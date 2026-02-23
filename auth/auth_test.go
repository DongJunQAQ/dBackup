package auth

import (
	"fmt"
	"testing"
)

var conf = Cert{
	AK: "HPUALK8VIW9T9AGYL7QM",
	SK: "Zz29pXObYTl4ynzcipz3ECdgmcljZAFKx5GOetzC",
}

//func TestSecret(t *testing.T) {
//	fmt.Println("随机生成的密钥([]byte类型):", genRandSecret())
//	fmt.Println("随机生成的密钥(String类型):", base64.StdEncoding.EncodeToString(genRandSecret()))
//}
//
//func TestEncrypt(t *testing.T) {
//	if len(currentSecret) != 32 {
//		t.Errorf("密钥长度不为32位")
//	}
//	data, _ := sonic.Marshal(conf)
//	res, err := encrypt(data, currentSecret)
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(res)
//}

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
