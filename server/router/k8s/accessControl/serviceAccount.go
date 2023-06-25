package accessControl

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ServiceAccountRouter struct{}

func (sar *ServiceAccountRouter) InitServiceAccountRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	serviceAccountRouter := Router.Group("serviceAccount")
	serviceAccountApi := api.ApiGroupApp.K8sApiGroup.AccessControlApi.ServiceAccountApi
	{
		serviceAccountRouter.POST("getServiceAccountList", serviceAccountApi.GetServiceAccountList)
		serviceAccountRouter.POST("getServiceAccountDetail", serviceAccountApi.GetServiceAccountDetail)
		serviceAccountRouter.POST("deleteServiceAccount", serviceAccountApi.DeleteServiceAccount)
	}
	return serviceAccountRouter
}
