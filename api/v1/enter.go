package v1

import (
	"main/api/v1/mall_api"
	"main/api/v1/manage_api"
)

type ApiGroup struct {
	ManageApiGroup manage_api.ManageGroup
	MallApiGroup   mall_api.MallGroup
}

var ApiGroupApp = new(ApiGroup)
