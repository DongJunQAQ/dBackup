package vault

import (
	dbakauth "dBackup/auth"
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

func CbrAuth() *cbr.CbrClient { //云备份服务的身份验证信息
	ak, sk, err := dbakauth.LoadAkSk() //从文件中读取ak和sk
	if err != nil {
		panic(err)
	}
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
