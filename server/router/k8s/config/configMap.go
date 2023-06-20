package config

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ConfigMapRouter struct{}

func (cr *ConfigMapRouter) InitConfigMapRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	configMapRouter := Router.Group("configMap")
	configMapApi := api.ApiGroupApp.K8sApiGroup.ConfigApi.ConfigMapApi
	{
		configMapRouter.POST("getConfigMapList", configMapApi.GetConfigMapList)
		configMapRouter.POST("getConfigMapDetail", configMapApi.GetConfigMapDetail)
		configMapRouter.POST("deleteConfigMap", configMapApi.DeleteConfigMap)
	}
	return configMapRouter
}
