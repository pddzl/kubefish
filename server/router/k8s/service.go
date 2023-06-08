package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ServiceRouter struct{}

func (sr *ServiceRouter) InitServiceRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	serviceRouter := Router.Group("service")
	serviceApi := api.ApiGroupApp.K8sApiGroup.ServiceApi
	{
		serviceRouter.POST("getServices", serviceApi.GetServices)
		serviceRouter.POST("getServiceDetail", serviceApi.GetServiceDetail)
		serviceRouter.POST("getServicePods", serviceApi.GetServicePods)
		serviceRouter.POST("deleteService", serviceApi.DeleteService)
	}
	return serviceRouter
}
