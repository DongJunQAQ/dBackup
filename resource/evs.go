package resource

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
	"github.com/pterm/pterm"
	"strings"
)

func ListEvs() {
	client := EvsAuth()
	request := &model.ListVolumesRequest{}
	limitRequest := int32(1000)
	request.Limit = &limitRequest
	response, err := client.ListVolumes(request)
	if err != nil {
		pterm.Error.Printf("查询失败: %v\n", err)
		return
	}
	resTableData := pterm.TableData{
		{"云硬盘ID", "云硬盘名称", "容量(GB)", "是否为启动盘", "挂载服务器ID"}, //初始化表格数据，第一行默认为表头
	}
	if response.Volumes != nil {
		for _, volume := range *response.Volumes {
			var resIdsShort []string
			if volume.Attachments != nil && len(volume.Attachments) > 0 {
				for _, res := range volume.Attachments {
					id := res.ServerId
					if len(id) > 8 { // 如果ID长度大于8，则只截取前8位，后面的加上省略号
						id = id[:8] + "..."
					}
					resIdsShort = append(resIdsShort, id)
				}
			}
			resIdsDisplay := strings.Join(resIdsShort, ", ")
			if resIdsDisplay == "" {
				resIdsDisplay = "未进行挂载"
			}
			resTableData = append(resTableData, []string{
				volume.Id,
				volume.Name,
				fmt.Sprintf("%d", volume.Size),
				volume.Bootable,
				resIdsDisplay,
			})
		}
	}
	_ = pterm.DefaultTable.
		WithHasHeader().
		WithBoxed().
		WithData(resTableData).
		Render()
	pterm.Success.Printf("共查询到 %d 个EVS云硬盘\n", len(*response.Volumes)) //打印查询成功后的输出
}
