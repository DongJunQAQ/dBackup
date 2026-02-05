package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/pterm/pterm"
	"os"
	"path/filepath"
)

type Cert struct { //凭证信息结构体
	AK string `json:"ak"`
	SK string `json:"sk"`
}

// var secret = getSecret() //动态获取密钥
var secret = []byte("a-very-secret-key-32-characters-") //静态密钥

func SaveAkSk(ak string, sk string) { //保存ak和sk至本地文件
	conf := Cert{
		AK: ak,
		SK: sk,
	}
	//1.构建文件路径
	homeDir, err := os.UserHomeDir() //获取家目录
	if err != nil {
		fmt.Println("无法获取家目录:", err)
		return
	}
	filePath := filepath.Join(homeDir, ".dbackup_config.json") //构建完整路径
	//2.序列化ak/sk结构体
	content, _ := sonic.Marshal(conf) //将结构体序列化为[]byte类型
	//3.加密ak/sk结构体
	encryptContent, err := encrypt(content, secret)
	if err != nil {
		fmt.Println("加密密钥失败:", err)
	}
	//4.封装为写入文件的JSON格式: {"auth": "..."}
	res := map[string]string{
		"auth": encryptContent,
	}
	resJson, _ := sonic.Marshal(res)
	//5.将加密后的内容写入文件
	err = os.WriteFile(filePath, resJson, 0600) //写入文件
	if err != nil {
		fmt.Println("保存文件失败:", err)
		return
	}
	pterm.Success.Printf("AK/SK已安全保存至: %s\n", filePath)
}

func LoadAkSk() (string, string, error) { //从本地文件中读取加密后的ak/sk并解密
	// 1. 获取家目录并构建路径
	homeDir, _ := os.UserHomeDir()
	filePath := filepath.Join(homeDir, ".dbackup_config.json")
	// 2. 读取文件内容
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", fmt.Errorf("读取文件失败: %v", err)
	}
	// 3. 解析外层 JSON (获取 {"auth": "..."} 中的字符串)
	var wrapper map[string]string
	if err := sonic.Unmarshal(fileData, &wrapper); err != nil {
		return "", "", fmt.Errorf("文件格式解析失败: %v", err)
	}
	encryptedBase64, _ := wrapper["auth"]
	// 4. Base64 解码密文
	ciphertext, _ := base64.StdEncoding.DecodeString(encryptedBase64)
	// 5. 初始化 AES-GCM 解密器
	block, _ := aes.NewCipher(secret)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", "", fmt.Errorf("密文数据过短")
	}
	// 6. 拆分 Nonce 和 密文，并执行解密
	nonce, actualCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, actualCiphertext, nil)
	if err != nil {
		return "", "", fmt.Errorf("解密校验失败（密钥不匹配或数据损坏）")
	}
	// 7. 将解密后的明文转为 Config 结构体
	var aksk Cert
	_ = sonic.Unmarshal(plaintext, &aksk)
	return aksk.AK, aksk.SK, nil
}
