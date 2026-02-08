package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func genRandSecret() []byte { //通过真随机的方式动态生成密钥
	randSecret := make([]byte, 32)                                  //由于密钥长度为32字节，因此需要32字节的空间
	if _, err := io.ReadFull(rand.Reader, randSecret); err != nil { //生成32字节的随机数
		panic(err)
	}
	return randSecret
}

func encrypt(data []byte, key []byte) (string, error) { //使用AES-GCM算法加密ak和sk
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil) //加密并附加nonce
	return base64.StdEncoding.EncodeToString(ciphertext), nil
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
