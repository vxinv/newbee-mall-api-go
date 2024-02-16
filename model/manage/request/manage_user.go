package request

import (
	"main/model/common/request"
	"main/model/manage"
)

type MallUserSearch struct {
	manage.MallUser
	request.PageInfo
}
