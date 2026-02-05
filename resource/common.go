package resource

import (
	dbakauth "dBackup/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
)

func EesAuth() *ecs.EcsClient {
	ak, sk, err := dbakauth.LoadAkSk() //从文件中读取ak和sk
	if err != nil {
		panic(err)
	}
	auth, _ := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		SafeBuild()
	reg, _ := region.SafeValueOf("cn-east-3") //华东-上海一
	hcClient, _ := ecs.EcsClientBuilder().
		WithRegion(reg).
		WithCredential(auth).
		SafeBuild()
	client := ecs.NewEcsClient(hcClient)
	return client
}
