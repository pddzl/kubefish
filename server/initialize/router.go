package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/middleware"
	"github.com/pddzl/kubefish/server/middleware/log"
)

func Routers() *gin.Engine {
	if global.KF_CONFIG.System.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	Router := gin.New()
	Router.Use(log.GinLogger(), log.GinRecovery(global.KF_CONFIG.System.Stack))

	// 跨域，如需跨域可以打开下面的注释
	// global.GVA_LOG.Info("use middleware cors")
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	global.KF_LOG.Info("register swagger handler")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 系统路由组
	systemRouter := router.RouterGroupApp.SystemRouterGroup

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	// 需要认证的路由
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitRoleRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitApiRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PublicGroup)
		systemRouter.InitJwtRouter(PublicGroup)
	}

	// k8s路由
	k8sRouter := router.RouterGroupApp.K8sRouterGroup

	K8sGroup := Router.Group("/k8s")
	K8sGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		k8sRouter.InitNodeRouter(K8sGroup)
		k8sRouter.InitResourceRouter(K8sGroup)
		k8sRouter.InitNamespaceRouter(K8sGroup)
		k8sRouter.InitPodRouter(K8sGroup)
		k8sRouter.InitReplicaSetRouter(K8sGroup)
		k8sRouter.InitDeploymentRouter(K8sGroup)
		k8sRouter.InitDaemonSetRouter(K8sGroup)
		k8sRouter.InitServiceRouter(K8sGroup)
	}

	global.KF_LOG.Info("router register success")
	return Router
}
