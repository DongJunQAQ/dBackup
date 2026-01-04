package vault

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	cbr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/region"
)

func CreateVault(name string, size int32) {
	ak := "HPUALK8VIW9T9AGYL7QM"
	sk := "Zz29pXObYTl4ynzcipz3ECdgmcljZAFKx5GOetzC"

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := cbr.NewCbrClient(
		cbr.CbrClientBuilder().
			WithRegion(region.ValueOf("cn-east-3")).
			WithCredential(auth).
			Build())

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
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
