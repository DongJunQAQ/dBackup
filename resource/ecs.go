package resource

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/pterm/pterm"
)

func ListEcs() {
	client := EesAuth()
	request := &model.NovaListServersRequest{}
	limitRequest := int32(25)
	request.Limit = &limitRequest
	response, err := client.NovaListServers(request)
	if err != nil {
		return
	}
	resTableData := pterm.TableData{
		{"服务器ID", "服务器名称"}, //初始化表格数据，第一行默认为表头
	}
	if response.Servers != nil { //遍历结果并将数据填充到tableData中
		for _, server := range *response.Servers {
			resTableData = append(resTableData, []string{ //将每一行数据加入resTableData中
				server.Id,
				server.Name,
			})
		}
	}
	_ = pterm.DefaultTable. //使用pterm渲染输出的结果表格
				WithHasHeader().        // 第一行作为表头
				WithBoxed().            // 加上外边框
				WithData(resTableData). // 注入数据
				Render()                // 渲染输出
	pterm.Success.Printf("共查询到 %d 台ECS服务器\n", len(*response.Servers)) //打印查询成功后的输出
}
