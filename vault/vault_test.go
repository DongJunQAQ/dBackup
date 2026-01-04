package vault

import "testing"

func TestListVault(t *testing.T) {
	ListVault()
}

func TestCreateVault(t *testing.T) {
	CreateVault("vault-test", 20)
	//如何捕获这里测试时的错误
}
