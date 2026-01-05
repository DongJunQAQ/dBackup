package vault

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

func CreateVault(name string, size int32) error { //创建存储库
	client := cbrAuth()
	request := &model.CreateVaultRequest{}
	chargingModeBilling := model.GetBillingCreateChargingModeEnum().POST_PAID
	isAutoRenewBilling := false
	isAutoPayBilling := false
	isMultiAzBilling := false
	billingVault := &model.BillingCreate{
		ConsistentLevel: model.GetBillingCreateConsistentLevelEnum().APP_CONSISTENT,
		ObjectType:      model.GetBillingCreateObjectTypeEnum().SERVER,
		ProtectType:     model.GetBillingCreateProtectTypeEnum().BACKUP,
		Size:            size,
		ChargingMode:    &chargingModeBilling,
		IsAutoRenew:     &isAutoRenewBilling,
		IsAutoPay:       &isAutoPayBilling,
		IsMultiAz:       &isMultiAzBilling,
	}
	thresholdVault := int32(80)
	smnNotifyVault := true
	demandBillingVault := false
	lockedVault := false
	vaultbody := &model.VaultCreate{
		Billing:       billingVault,
		Name:          name,
		Resources:     []model.ResourceCreate{},
		Threshold:     &thresholdVault,
		SmnNotify:     &smnNotifyVault,
		DemandBilling: &demandBillingVault,
		Locked:        &lockedVault,
	}
	request.Body = &model.VaultCreateReq{
		Vault: vaultbody,
	}
	response, err := client.CreateVault(request)
	if err == nil {
		if response.Vault != nil {
			fmt.Println(response.Vault.Id)
		}
		return nil
	} else {
		return err
	}
}
