package vault

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

func AddResources(vaultId string, resourcesID string, resourcesType string) error { //往存储库中添加资源
	convertedType := ""    //转换后的类型
	switch resourcesType { //资源类型转换
	case "server":
		convertedType = "OS::Nova::Server"
	case "volume":
		convertedType = "OS::Cinder::Volume"
	default:
		return fmt.Errorf("未知类型")
	}
	client := cbrAuth()
	request := &model.AddVaultResourceRequest{}
	request.VaultId = vaultId
	var listResourcesbody = []model.ResourceCreate{
		{
			Id:   resourcesID,
			Type: convertedType,
		},
	}
	request.Body = &model.VaultAddResourceReq{
		Resources: listResourcesbody,
	}
	response, err := client.AddVaultResource(request)
	if err == nil {
		for _, resource := range *response.AddResourceIds {
			fmt.Println(resource)
		}
		return nil
	} else {
		return err
	}
}
