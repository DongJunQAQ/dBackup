package vault

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	cbr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/region"
)

func listVault() { //列出备份存储库
	//ak := os.Getenv("CLOUD_SDK_AK")
	ak := "HPUALK8VIW9T9AGYL7QM"
	//sk := os.Getenv("CLOUD_SDK_SK")
	sk := "Zz29pXObYTl4ynzcipz3ECdgmcljZAFKx5GOetzC"

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := cbr.NewCbrClient(
		cbr.CbrClientBuilder().
			WithRegion(region.ValueOf("cn-east-3")).
			WithCredential(auth).
			Build())

	request := &model.ListVaultRequest{}
	response, err := client.ListVault(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
