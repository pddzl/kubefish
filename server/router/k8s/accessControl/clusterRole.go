package accessControl

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ClusterRoleRouter struct{}

func (rr *ClusterRoleRouter) InitClusterRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	clusterRoleRouter := Router.Group("clusterRole")
	clusterRoleApi := api.ApiGroupApp.K8sApiGroup.AccessControlApi.ClusterRoleApi
	{
		clusterRoleRouter.POST("getClusterRoleList", clusterRoleApi.GetClusterRoleList)
		clusterRoleRouter.POST("getClusterRoleDetail", clusterRoleApi.GetClusterRoleDetail)
		clusterRoleRouter.POST("deleteClusterRole", clusterRoleApi.DeleteClusterRole)
	}
	return clusterRoleRouter
}
