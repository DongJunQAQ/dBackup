package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func genTrueRand(size int) []byte { //根据指定的字节大小生成真随机数，如传入32表示生成32字节的真随机数
	trueRand := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, trueRand); err != nil {
		panic(err)
	}
	return trueRand
}

func encrypt(plaintext, secret []byte) (string, error) { //使用AES-GCM算法加密ak和sk
	block, err := aes.NewCipher(secret) //根据提供的密钥创建一个AES加密器实例，因为需要使用密钥长度为256位的AES-256算法，所以只需要传入32字节的密钥，Go就会自动选择AES-256算法(32*8=256)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block) //将基础的AES加密器实例升级为GCM模式(Galois Counter Mode，伽罗瓦计数器模式)，这一步让加密具备了“认证”功能。它能生成一个认证标签（Tag），在解密时校验数据是否被动过
	if err != nil {
		return "", err
	}
	nonce := genTrueRand(gcm.NonceSize()) //生成随机Nonce(临时随机数，Number used once)，NonceSize()通常是12字节，对于同一个密钥，绝对不能使用相同的Nonce加密两段不同的数据
	//因此这里使用真随机数生成器来填充Nonce，使得每次的Nonce都不一样，因为Nonce是真随机生成的，所以即使用同样的密钥加密同样的明文两次，得到的密文也会完全不同。这能有效防止重放攻击
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil) //实际的加密操作
	//Seal()函数的参数详解：
	//dst: 目标切片，加密后的结果（密文+Tag）会追加到这个切片后面，如果传入nonce则最终结果为“[12字节Nonce]+[加密后的内容]+[16字节Tag]”
	//而如果传入的是plaintext[:0]即原地覆盖则最终结果为“[加密后的内容]+[16字节Tag]”
	//nonce: 临时随机数
	//plaintext: 明文
	//additionalData: 附加数据，只用于认证不会被加密的数据，假设在传输一个数据包时希望“用户ID”不加密，但又不允许被篡改，则将其设置为附加数据传入，而在解密时必须传入相同的"附加数据"内容才能解密成功
	return base64.StdEncoding.EncodeToString(ciphertext), nil //由于加密后的密文是二进制字节流，所以这里需要将其转换为Base64字符串格式，以便写入配置文件中
}

func decrypt(ciphertext, realSecret []byte) ([]byte, error) { //解密
	block, blockErr := aes.NewCipher(realSecret)
	if blockErr != nil {
		return []byte(""), fmt.Errorf("secret可能已被篡改")
	}
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return []byte(""), fmt.Errorf("密文数据过短")
	}
	nonce, actualCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, actualCiphertext, nil)
	if err != nil {
		return []byte(""), fmt.Errorf("解密校验失败（密钥不匹配或数据损坏）")
	}
	return plaintext, nil
}
