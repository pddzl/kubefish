package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ResourceRouter struct{}

func (rr *ResourceRouter) InitResourceRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	resourceRouter := Router.Group("resource")
	resourceApi := api.ApiGroupApp.K8sApiGroup.ResourceApi.ResourceApi
	{
		resourceRouter.POST("getResourceRaw", resourceApi.GetResourceRaw)
	}
	return resourceRouter
}
