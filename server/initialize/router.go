package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/middleware"
	"github.com/pddzl/kubefish/server/middleware/log"
	"github.com/pddzl/kubefish/server/router"
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
		// cluster
		k8sRouter.ClusterRouter.InitNodeRouter(K8sGroup)
		k8sRouter.ClusterRouter.InitNamespaceRouter(K8sGroup)
		// workloads
		k8sRouter.WorkloadsRouter.InitPodRouter(K8sGroup)
		k8sRouter.WorkloadsRouter.InitReplicaSetRouter(K8sGroup)
		k8sRouter.WorkloadsRouter.InitDeploymentRouter(K8sGroup)
		k8sRouter.WorkloadsRouter.InitDaemonSetRouter(K8sGroup)
		// service
		k8sRouter.ServiceRouter.InitServicesRouter(K8sGroup)
		k8sRouter.ServiceRouter.InitIngressRouter(K8sGroup)
		// config
		k8sRouter.ConfigRouter.InitConfigMapRouter(K8sGroup)
		k8sRouter.ConfigRouter.InitSecretRouter(K8sGroup)
		// accessControl
		k8sRouter.AccessControlRouter.InitServiceAccountRouter(K8sGroup)
		k8sRouter.AccessControlRouter.InitRoleRouter(K8sGroup)
		k8sRouter.AccessControlRouter.InitRoleBindingRouter(K8sGroup)
		k8sRouter.AccessControlRouter.InitClusterRoleRouter(K8sGroup)
		k8sRouter.AccessControlRouter.InitClusterRoleBindingRouter(K8sGroup)
		// resource
		k8sRouter.ResourceRouter.InitResourceRouter(K8sGroup)
	}

	global.KF_LOG.Info("router register success")
	return Router
}
