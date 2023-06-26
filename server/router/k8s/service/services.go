package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ServicesRouter struct{}

func (sr *ServicesRouter) InitServicesRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	serviceRouter := Router.Group("service")
	serviceApi := api.ApiGroupApp.K8sApiGroup.ServiceApi.ServicesApi
	{
		serviceRouter.POST("getServices", serviceApi.GetServices)
		serviceRouter.POST("getServiceDetail", serviceApi.GetServiceDetail)
		serviceRouter.POST("getServicePods", serviceApi.GetServicePods)
		serviceRouter.POST("deleteService", serviceApi.DeleteService)
	}
	return serviceRouter
}
