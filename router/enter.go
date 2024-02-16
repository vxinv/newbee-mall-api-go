package router

import (
	"main/router/mall"
	"main/router/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	Mall   mall.MallRouterGroup
}

var RouterGroupApp = new(RouterGroup)
