package vault

import (
	"fmt"
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
		//返回存储库ID不存在的错误
		return fmt.Errorf("%w", ErrNotExists)
	}
	return err
}
