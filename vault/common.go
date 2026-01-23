package vault

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	cbr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/region"
)

var (
	ErrUnknownType = errors.New("未知类型")
	ErrNotExists   = errors.New("存储库或资源不存在")
	ErrFormat      = errors.New("ID格式有误")
)

func cbrAuth() *cbr.CbrClient { //云备份服务的身份验证信息
	//ak := os.Getenv("CLOUD_SDK_AK")
	ak := "HPUALK8VIW9T9AGYL7QM"
	//sk := os.Getenv("CLOUD_SDK_SK")
	sk := "Zz29pXObYTl4ynzcipz3ECdgmcljZAFKx5GOetzC"
	auth, _ := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		SafeBuild()
	reg, _ := region.SafeValueOf("cn-east-3") //华东-上海一
	hcClient, _ := cbr.CbrClientBuilder().
		WithRegion(reg).
		WithCredential(auth).
		SafeBuild()
	client := cbr.NewCbrClient(hcClient)
	return client
}
