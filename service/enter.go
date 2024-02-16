package service

import (
	"main/service/example"
	"main/service/mall_service"
	"main/service/manage_service"
)

type ServiceGroup struct {
	ExampleServiceGroup example.ServiceGroup
	ManageServiceGroup  manage_service.ManageServiceGroup
	MallServiceGroup    mall_service.MallServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
