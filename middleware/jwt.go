package middleware

import (
	"github.com/gin-gonic/gin"
	"main/model/common/response"
	"main/service/mall_service"
	"main/service/manage_service"
	"time"
)

var (
	ms  manage_service.ManageAdminUserTokenService
	mus mall_service.MallUserTokenService
)

//var manageAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserTokenService
//var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService

func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailed(nil, "未登录或非法访问", c)
			c.Abort()
			return
		}
		err, mallAdminUserToken := ms.ExistAdminToken(token)
		if err != nil {
			response.FailWithDetailed(nil, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if time.Now().After(mallAdminUserToken.ExpireTime) {
			response.FailWithDetailed(nil, "授权已过期", c)
			err = ms.DeleteMallAdminUserToken(token)
			if err != nil {
				return
			}
			c.Abort()
			return
		}
		c.Next()
	}

}

func UserJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.UnLogin(nil, c)
			c.Abort()
			return
		}
		err, _ := mus.ExistUserToken(token)
		if err != nil {
			response.UnLogin(nil, c)
			c.Abort()
			return
		}
		if time.Now().After(mus.ExpireTime) {
			response.FailWithDetailed(nil, "授权已过期", c)
			err = mus.DeleteMallUserToken(token)
			if err != nil {
				return
			}
			c.Abort()
			return
		}
		c.Next()
	}

}
