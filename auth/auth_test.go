package auth

import (
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

var conf = Cert{
	AK: "HPUALK8VIW9T9AGYL7QM",
	SK: "Zz29pXObYTl4ynzcipz3ECdgmcljZAFKx5GOetzC",
}

func TestEncrypt(t *testing.T) {
	if len(secret) != 32 {
		t.Errorf("密钥长度不为32位")
	}
	data, _ := sonic.Marshal(conf)
	res, err := encrypt(data, secret)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}

func TestSaveAkSk(t *testing.T) {
	ak := "HPUALK8VIW9T9AGYL7QM"
	sk := "Zz29pXObYTl4ynzcipz3ECdgmcljZAFKx5GOetzC"
	SaveAkSk(ak, sk)
}

func TestLoadAkSk(t *testing.T) {
	ak, sk, err := LoadAkSk()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("AK:", ak)
	fmt.Println("SK:", sk)
}
