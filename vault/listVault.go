package vault

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	cbr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/region"
)

func ListVault() { //列出备份存储库
	//ak := os.Getenv("CLOUD_SDK_AK")
	ak := "HPUALK8VIW9T9AGYL7QM"
	//sk := os.Getenv("CLOUD_SDK_SK")
	sk := "Zz29pXObYTl4ynzcipz3ECdgmcljZAFKx5GOetzC"

	auth, _ := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		SafeBuild()

	reg, _ := region.SafeValueOf("cn-east-3")
	hcClient, _ := cbr.CbrClientBuilder().
		WithRegion(reg).
		WithCredential(auth).
		SafeBuild()
	client := cbr.NewCbrClient(hcClient)

	request := &model.ListVaultRequest{}
	response, err := client.ListVault(request)
	if err == nil {
		fmt.Printf("%-36s | %-15s | %-5s | %-s\n", "ID", "Name", "Size", "Resources")
		fmt.Println("--------------------------------------------------------------------------------")
		// 检查 Vaults 是否为空
		if response.Vaults != nil {
			for _, v := range *response.Vaults {
				size := 0
				if v.Billing != nil {
					size = int(v.Billing.Size)
				}
				fmt.Printf("%-36s | %-15s | %-5d | %v\n", v.Id, v.Name, size, v.Resources)
			}
		}
	} else {
		fmt.Println(err)
	}
}
