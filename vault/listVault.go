package vault

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

func ListVault() { //列出备份存储库
	client := cbrAuth()
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
