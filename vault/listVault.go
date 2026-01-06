package vault

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

func ListVault() { //列出备份存储库
	client := cbrAuth()
	request := &model.ListVaultRequest{}
	response, err := client.ListVault(request)
	if err == nil {
		fmt.Printf("%-36s | %-15s | %-5s | %-s\n", "ID", "Name", "Size", "Resources")
		fmt.Println("--------------------------------------------------------------------------------")
		if response.Vaults != nil {
			for _, v := range *response.Vaults {
				size := 0
				if v.Billing != nil {
					size = int(v.Billing.Size)
				}
				//提取资源ID并组成切片
				var resIds []string
				if v.Resources != nil {
					for _, res := range v.Resources {
						resIds = append(resIds, res.Id)
					}
				}
				resIdsBytes, _ := sonic.Marshal(resIds) //序列化资源ID切片，使其以JSON格式展示
				fmt.Printf("%-36s | %-15s | %-5d | %v\n", v.Id, v.Name, size, string(resIdsBytes))
			}
		}
	} else {
		fmt.Println(err)
	}
}
