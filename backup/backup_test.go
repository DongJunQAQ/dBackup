package backup

import (
	"dBackup/vault"
	"errors"
	"fmt"
	"testing"
)

func TestCreateCheckpoint(t *testing.T) {
	testCases := []struct {
		name        string
		vaultId     string
		expectedErr error
	}{
		{"ExistsVaultCreateCheckpoint", "e195589e-7222-41ac-a28a-b92037a8bf68", nil},   //对已存在的存储库创建检查点
		{"NotExistsVault", "88888888-8888-8888-8888-888888888888", vault.ErrNotExists}, //对不存在的存储库创建检查点
		{"ErrFormat", "88888888-8888-8888-8888-8888888888889", vault.ErrFormat},        //为ID格式错误的存储库创建检查点
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := CreateCheckpoint(tc.vaultId)
			if tc.expectedErr != nil {
				if !errors.Is(err, tc.expectedErr) { //如果预期错误与实际错误不符
					fmt.Println("错误类型: ", err)
					t.Errorf("预期错误为: %q, 实际错误为: %q", tc.expectedErr, err)
				}
				return //提前结束测试
			}
			if err != nil {
				t.Fatalf("预期执行备份成功无报错但实际报错: %q", err)
			}
		})
	}
}
