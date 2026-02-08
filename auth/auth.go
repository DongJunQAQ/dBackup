package auth

import (
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

var currentSecret = genRandSecret()                              //动态生成[]byte类型的密钥
var secretStr = base64.StdEncoding.EncodeToString(currentSecret) //将密钥转换为String类型

func getConfPath() (string, error) { //获取dBackup配置文件的完整路径
	homeDir, err := os.UserHomeDir() //获取家目录
	if err != nil {
		return "", fmt.Errorf("无法获取家目录: %v\n", err)

	}
	filePath := filepath.Join(homeDir, ".dbackup_config.json") //拼接完整路径
	return filePath, nil
}

func SaveAkSk(ak string, sk string) { //保存ak和sk至本地文件
	conf := Cert{
		AK: ak,
		SK: sk,
	}
	filePath, err := getConfPath()
	if err != nil {
		fmt.Println("无法获取配置文件的路径:", err)
	}
	//2.序列化ak/sk结构体
	content, _ := sonic.Marshal(conf) //将结构体序列化为[]byte类型
	//3.加密ak/sk结构体
	encryptContent, err := encrypt(content, currentSecret)
	if err != nil {
		fmt.Println("加密密钥失败:", err)
	}
	//4.封装为写入文件的JSON格式: {"secret": "...","auth": "..."}
	res := map[string]string{
		"secret": secretStr,
		"auth":   encryptContent,
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
	filePath, err := getConfPath()
	if err != nil {
		return "", "", fmt.Errorf("无法获取配置文件的路径: %v", err)
	}
	// 2. 读取文件内容
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", fmt.Errorf("读取配置文件失败(请先使用login子命令登录): %v", err)
	}
	// 3. 解析外层 JSON (获取 {"auth": "..."} 中的字符串)
	var wrapper map[string]string
	if err := sonic.Unmarshal(fileData, &wrapper); err != nil {
		return "", "", fmt.Errorf("文件格式解析失败: %v", err)
	}
	authBase64 := wrapper["auth"]
	//标准的Base64编码是4字符一组，如果在合法的Base64字符串末尾加上不足4字符的内容，
	//Base64解码器在处理时，如果这多出的字符不足以构成一个新的有效字节，
	//解码器可能会忽略末尾不完整的位，也就使得在密文后面添加了不足4字符的内容后会导致密文依旧有效，
	//因此需要添加Base64内容是否为4的倍数的校验
	if len(authBase64)%4 != 0 {
		return "", "", fmt.Errorf("auth长度错误")
	}
	// 4. Base64 解码密文
	ciphertext, _ := base64.StdEncoding.DecodeString(authBase64)
	// 5. 初始化 AES-GCM 解密器
	secretBase64 := wrapper["secret"]
	if len(secretBase64)%4 != 0 {
		return "", "", fmt.Errorf("secret长度错误")
	}
	realSecret, _ := base64.StdEncoding.DecodeString(secretBase64) //从配置文件中读取secret字段以获取密钥
	//6.解密密钥
	plaintext, err := decrypt(ciphertext, realSecret)
	if err != nil {
		return "", "", fmt.Errorf("解密密钥失败: %v", err)
	}
	// 7. 将解密后的明文转为 Config 结构体
	var aksk Cert
	_ = sonic.Unmarshal(plaintext, &aksk)
	return aksk.AK, aksk.SK, nil
}
