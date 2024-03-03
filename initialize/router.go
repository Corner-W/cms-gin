package initialize

import (
	"cms/router"
	"fmt"

	
	"cms/common/enums"

	"go.uber.org/zap"

	"cms/global"
	"cms/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "cms/docs"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.GinRecovery(true)) //记录Panic日志


	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	// swagger docs 设置
	if global.Env != enums.PROD {
		global.GVA_LOG.Info("set swagger path:env=========" + global.Env.String())
		Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		global.GVA_LOG.Info("register swagger handler")
	}

	// 健康监测
	Router.GET("/health", func(c *gin.Context) {
		if global.GVA_DB != nil && global.GVA_LOG != nil {
			sqlDB, _ := global.GVA_DB.DB()
			dbStats := sqlDB.Stats()
			global.GVA_LOG.Info("", zap.String("info", fmt.Sprintf("sql.DB statistics: %+v\n", dbStats)))
			//fmt.Printf("sql.DB statistics: %+v\n", dbStats)
		}
		c.JSON(200, "ok")
	})

	// 方便统一添加路由组前缀 多服务器上线使用
	//获取路由组实例
	cmsRouter := router.GroupApp.Cms.CmsRouter



	baseGroupV1 := Router.Group("/api/v1")

	{
		cmsRouter.InitRouter(baseGroupV1)


	}
	global.GVA_LOG.Info("router register success")
	return Router
}
