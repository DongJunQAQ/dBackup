package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func getSecret() []byte { //通过真随机的方式获取密钥
	secret := make([]byte, 32) //由于密钥长度为32字节，因此需要生成32字节的随机数
	if _, err := io.ReadFull(rand.Reader, secret); err != nil {
		panic(err)
	}
	return secret
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
