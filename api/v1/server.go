package v1

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"main/global"
	"main/middleware"
	"main/router"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	//商城后管路由
	manageRouter := router.RouterGroupApp.Manage
	ManageGroup := Router.Group("manage_api-api")
	PublicGroup := Router.Group("")

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		//商城后管路由初始化
		manageRouter.InitManageAdminUserRouter(ManageGroup)
		manageRouter.InitManageGoodsCategoryRouter(ManageGroup)
		manageRouter.InitManageGoodsInfoRouter(ManageGroup)
		manageRouter.InitManageCarouselRouter(ManageGroup)
		manageRouter.InitManageIndexConfigRouter(ManageGroup)
		manageRouter.InitManageOrderRouter(ManageGroup)
	}
	//商城前端路由
	mallRouter := router.RouterGroupApp.Mall
	MallGroup := Router.Group("api")
	{
		// 商城前端路由
		mallRouter.InitMallCarouselIndexRouter(MallGroup)
		mallRouter.InitMallGoodsInfoIndexRouter(MallGroup)
		mallRouter.InitMallGoodsCategoryIndexRouter(MallGroup)
		mallRouter.InitMallUserRouter(MallGroup)
		mallRouter.InitMallUserAddressRouter(MallGroup)
		mallRouter.InitMallShopCartRouter(MallGroup)
		mallRouter.InitMallOrderRouter(MallGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
