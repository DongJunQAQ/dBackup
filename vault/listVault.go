package vault

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
	"github.com/pterm/pterm"
	"strings"
)

func ListVault() { //列出备份存储库
	client := CbrAuth()
	request := &model.ListVaultRequest{}
	response, err := client.ListVault(request)
	if err != nil {
		return
	}
	tableData := pterm.TableData{
		{"存储库ID", "存储库名称", "容量(GB)", "关联资源数量", "资源ID列表"},
	}
	if response.Vaults != nil {
		for _, v := range *response.Vaults {
			sizeStr := "0"
			if v.Billing != nil {
				sizeStr = pterm.Sprint(v.Billing.Size)
			}
			var resIdsShort []string
			if v.Resources != nil {
				for _, res := range v.Resources {
					id := res.Id
					if len(id) > 8 { // 如果ID长度大于8，则只截取前8位，后面的加上省略号
						id = id[:8] + "..."
					}
					resIdsShort = append(resIdsShort, id)
				}
			}
			resIdsDisplay := strings.Join(resIdsShort, ", ")
			if resIdsDisplay == "" {
				resIdsDisplay = "无关联资源"
			}
			resCount := pterm.Sprint(len(resIdsShort))
			tableData = append(tableData, []string{
				v.Id,
				v.Name,
				sizeStr,
				resCount,
				resIdsDisplay,
			})
		}
	}
	_ = pterm.DefaultTable.
		WithHasHeader().
		WithBoxed().
		WithHeaderStyle(pterm.NewStyle(pterm.FgLightCyan, pterm.Bold)).
		WithData(tableData).
		Render()
	pterm.Success.Printf("共查询到 %d 个CBR存储库\n", len(*response.Vaults))
}
