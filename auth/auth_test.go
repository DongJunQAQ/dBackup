package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

var conf = Cert{
	AK: "HPUALK8VIW9T9AGYL7QM",
	SK: "Zz29pXObYTl4ynzcipz3ECdgmcljZAFKx5GOetzC",
}

func TestGenSecret(t *testing.T) {
	secret := genTrueRand(32)
	if len(secret) != 32 {
		t.Fatalf("secret长度应为32")
	}
	fmt.Println("随机生成的密钥([]byte类型):", secret)
	fmt.Println("随机生成的密钥(String类型):", base64.StdEncoding.EncodeToString(genTrueRand(32)))
}

func TestEncrypt(t *testing.T) {
	data, _ := sonic.Marshal(conf) //将go对象序列化为字节流格式
	res, err := encrypt(data, genTrueRand(32))
	if err != nil {
		t.Error("加密失败: ", err)
	}
	fmt.Println("auth字段的值为: ", res)
}

func TestSaveAkSk(t *testing.T) {
	SaveAkSk(conf.AK, conf.SK)
}

func TestLoadAkSk(t *testing.T) {
	ak, sk, err := LoadAkSk()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("AK:", ak)
	fmt.Println("SK:", sk)
}
