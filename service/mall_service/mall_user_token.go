package mall_service

import (
	"main/global"
	"main/model/mall"
	"time"
)

type MallUserTokenService struct {
	ExpireTime time.Time
}

func (m *MallUserTokenService) ExistUserToken(token string) (err error, mallUserToken mall.MallUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallUserToken).Error
	return
}

func (m *MallUserTokenService) DeleteMallUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]mall.MallUserToken{}, "token =?", token).Error
	return err
}
