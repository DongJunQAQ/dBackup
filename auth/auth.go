package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/pterm/pterm"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Cert struct { //凭证信息结构体
	AK string `json:"ak"`
	SK string `json:"sk"`
}

var secretByte = genTrueRand(32)                              //动态生成[]byte类型的密钥
var secretStr = base64.StdEncoding.EncodeToString(secretByte) //将密钥转换为String类型

func getConfPath() (string, error) { //获取dBackup配置文件的完整路径
	homeDir, err := os.UserHomeDir() //获取家目录
	if err != nil {
		return "", fmt.Errorf("无法获取家目录: %v\n", err)
	}
	filePath := filepath.Join(homeDir, ".dbackup_config.json") //拼接完整路径
	return filePath, nil
}

func getConfigFile() *viper.Viper { //获取配置文件中字段的内容
	config := viper.New()
	home, _ := os.UserHomeDir() //获取用户的家目录路径
	config.AddConfigPath(home)
	config.SetConfigType("json")
	config.SetConfigName(".dbackup_config")
	err := config.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError //文件不存在的错误
		if errors.As(err, &configFileNotFoundError) {             //判断此错误是否为“文件不存在”的错误
			pterm.Error.Println("配置文件不存在(请先使用login子命令登录)")
		} else {
			panic(fmt.Errorf("读取配置文件错误：%w", err))
		}
	}
	return config
}

func SaveAkSk(ak string, sk string) { //保存ak和sk至本地文件
	aksk := Cert{
		AK: ak,
		SK: sk,
	}
	content, _ := sonic.Marshal(aksk)         //将凭证结构体序列化为[]byte类型
	auth, err := encrypt(content, secretByte) //加密ak/sk结构体，得到auth字段中的内容
	if err != nil {
		fmt.Println("加密密钥失败:", err)
	}
	res := map[string]string{ //封装为写入文件的JSON格式: {"secret": "...","auth": "..."}
		"secret": secretStr,
		"auth":   auth,
	}
	resJson, _ := sonic.Marshal(res) //将Go的map对象序列化[]byte字节流类型
	filePath, err := getConfPath()   //获取配置文件路径
	if err != nil {
		fmt.Println("无法获取配置文件的路径:", err)
	}
	err = os.WriteFile(filePath, resJson, 0600) //将加密后的内容写入配置文件
	if err != nil {
		fmt.Println("保存文件失败:", err)
		return
	}
	pterm.Success.Printf("AK/SK已安全保存至: %s\n", filePath)
}

func LoadAkSk() (string, string, error) { //从本地文件中读取加密后的ak/sk并解密
	conf := getConfigFile()
	authBase64 := conf.GetString("auth") //以字符串格式获取auth字段中的内容(base64编码)
	secretBase64 := conf.GetString("secret")
	//标准的Base64编码是4字符一组，如果在合法的Base64字符串末尾加上不足4字符的内容，
	//Base64解码器在处理时，如果这多出的字符不足以构成一个新的有效字节，
	//解码器可能会忽略末尾不完整的位，也就使得在密文后面添加了不足4字符的内容后会导致密文依旧有效，
	//因此需要添加Base64内容是否为4的倍数的校验
	if len(authBase64)%4 != 0 {
		return "", "", fmt.Errorf("auth长度错误")
	}
	if len(secretBase64)%4 != 0 {
		return "", "", fmt.Errorf("secret长度错误")
	}
	ciphertext, _ := base64.StdEncoding.DecodeString(authBase64) //将base64字符串格式的auth解码为[]byte字节流格式
	secret, _ := base64.StdEncoding.DecodeString(secretBase64)   //将base64字符串格式的secret解码为[]byte字节流格式
	plaintext, err := decrypt(ciphertext, secret)                //使用密钥解密
	if err != nil {
		return "", "", fmt.Errorf("解密失败: %v", err)
	}
	var aksk Cert
	_ = sonic.Unmarshal(plaintext, &aksk) //将解密获得的[]byte字节流格式的明文反序列化为Go对象
	return aksk.AK, aksk.SK, nil
}
