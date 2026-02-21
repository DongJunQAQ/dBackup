package resource

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
)

func ListEvs() {
	client := EvsAuth()
	request := &model.ListVolumesRequest{}
	limitRequest := int32(1000)
	request.Limit = &limitRequest
	response, err := client.ListVolumes(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
