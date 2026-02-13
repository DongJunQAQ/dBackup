package backup

import (
	"dBackup/vault"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
	"github.com/pterm/pterm"
)

func CreateCheckpoint(vaultId string) error { //为存储库创建备份还原点，手动备份存储库中所有的资源
	client := vault.CbrAuth()
	request := &model.CreateCheckpointRequest{}
	checkpointbody := &model.VaultBackup{
		VaultId: vaultId,
	}
	request.Body = &model.VaultBackupReq{
		Checkpoint: checkpointbody,
	}
	response, err := client.CreateCheckpoint(request)
	if err == nil { //当创建备份点无报错
		if response.Checkpoint != nil { //且还原点信息不为空时，说明执行备份成功
			pterm.Success.Printf("存储库%s正在执行备份请稍后...\n", vaultId)
		}
	} else {
		var serviceErr *sdkerr.ServiceResponseError
		if errors.As(err, &serviceErr) {
			if serviceErr.ErrorCode == "BackupService.9900" {
				return fmt.Errorf("%w", vault.ErrFormat) //返回ID格式不对的错误
			}
			if serviceErr.ErrorCode == "BackupService.6302" || serviceErr.ErrorCode == "BackupService.6105" {
				return fmt.Errorf("%w", vault.ErrNotExists) //返回资源或存储库不存在的错误
			}
		}
	}
	return err
}
