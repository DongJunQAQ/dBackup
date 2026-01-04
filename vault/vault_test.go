package vault

import "testing"

func TestListVault(t *testing.T) {
	ListVault()
}

func TestCreateVault(t *testing.T) {
	err := CreateVault("vault-test", 20)
	if err != nil {
		t.Fatalf("测试失败: %v", err)
	}
}
