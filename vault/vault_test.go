package vault

import (
	"errors"
	"fmt"
	"testing"
)

func TestListVault(t *testing.T) {
	ListVault()
}

func TestCreateVault(t *testing.T) {
	err := CreateVault("vault-test", 20)
	if err != nil {
		t.Fatalf("创建存储库失败: %v", err)
	}
}

func TestAddResources(t *testing.T) {
	//使用表驱动测试
	testCases := []struct {
		name          string
		vaultId       string
		resourcesID   string
		resourcesType string
		expectedErr   error
	}{ //定义并初始化一个匿名结构体切片
		{"AddNew", "ad6910b3-6c77-45bb-8651-276adce80fb8", "0a28470b-7f52-4fc6-97a0-68dced426e08", "server", nil}, //需实时填写新资源的ID
		{"UnknownType", "77c1fc62-2db4-4ff8-87c4-019459901073", "3d562067-bb98-48b3-b1a1-3a325af075be", "server123", ErrUnknownType},
		{"NotExists", "88888888-8888-8888-8888-888888888888", "88888888-8888-8888-8888-888888888888", "server", ErrNotExists},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := AddResources(tc.vaultId, tc.resourcesID, tc.resourcesType)
			if tc.expectedErr != nil {
				if !errors.Is(err, tc.expectedErr) { //如果预期错误与实际错误不符
					fmt.Println("错误类型: ", err)
					t.Errorf("预期错误为: %q, 实际错误为: %q", tc.expectedErr, err)
				}
				return //提前结束测试
			}
			if err != nil {
				t.Fatalf("预期添加资源成功无报错但实际报错: %q", err)
			}
		})
	}
}
