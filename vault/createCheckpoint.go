package vault

import (
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

func CreateCheckpoint(vaultId string) error { //为存储库创建备份还原点，手动备份存储库中所有的资源
	client := cbrAuth()
	request := &model.CreateCheckpointRequest{}
	checkpointbody := &model.VaultBackup{
		VaultId: vaultId,
	}
	request.Body = &model.VaultBackupReq{
		Checkpoint: checkpointbody,
	}
	response, err := client.CreateCheckpoint(request)
	if err == nil {
		if response.Checkpoint != nil {
			fmt.Println(response.Checkpoint.Id) //打印还原点ID
			return nil
		}
	} else {
		var serviceErr *sdkerr.ServiceResponseError
		if errors.As(err, &serviceErr) {
			if serviceErr.ErrorCode == "BackupService.9900" {
				return fmt.Errorf("%w", ErrFormat) //返回ID格式不对的错误
			}
			if serviceErr.ErrorCode == "BackupService.6302" || serviceErr.ErrorCode == "BackupService.6105" {
				return fmt.Errorf("%w", ErrNotExists) //返回资源或存储库不存在的错误
			}
		}
	}
	return err
}
