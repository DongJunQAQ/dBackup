package vault

import (
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

var (
	ErrUnknownType = errors.New("未知类型")
	ErrNotExists   = errors.New("存储库或资源不存在")
	ErrFormat      = errors.New("ID格式有误")
)

func AddResources(vaultId string, resourcesID string, resourcesType string) error { //往存储库中添加资源
	convertedType := ""    //转换后的类型
	switch resourcesType { //资源类型转换
	case "server":
		convertedType = "OS::Nova::Server"
	case "volume":
		convertedType = "OS::Cinder::Volume"
	default:
		return fmt.Errorf("%w: %s", ErrUnknownType, resourcesType)
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
		var serviceErr *sdkerr.ServiceResponseError //华为云SDK专门定义的错误结构体，里面包含了StatusCode、RequestId等详细信息
		if errors.As(err, &serviceErr) {            //会检查err是不是属于*sdkerr.ServiceResponseError这种类型,如果是这种类型的错误errors.As会返回true，并自动把err里面的具体内容填充到serviceErr变量里
			if serviceErr.ErrorCode == "BackupService.9900" {
				return fmt.Errorf("%w", ErrFormat) //返回ID格式不对的错误
			}
			if serviceErr.ErrorCode == "BackupService.6302" || serviceErr.ErrorCode == "BackupService.6105" {
				return fmt.Errorf("%w", ErrNotExists) //返回资源或存储库不存在的错误
			}
		}
		return err //返回其他错误
	}
}
