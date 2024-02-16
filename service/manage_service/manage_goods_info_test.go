package manage_service

import (
	"main/initialize"
	"testing"
)

func TestManageGoodsInfoService_GetMallGoodsInfoInfoList(t *testing.T) {
	mysql := initialize.MockGormMysql()
	mysql.Raw("select 1")
}
