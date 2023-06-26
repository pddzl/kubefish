package accessControl

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ClusterRoleBindingRouter struct{}

func (rr *ClusterRoleBindingRouter) InitClusterRoleBindingRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	clusterRoleBindingRouter := Router.Group("clusterRoleBinding")
	clusterRoleBindingApi := api.ApiGroupApp.K8sApiGroup.AccessControlApi.ClusterRoleBindingApi
	{
		clusterRoleBindingRouter.POST("getClusterRoleBindingList", clusterRoleBindingApi.GetClusterRoleBindingList)
		clusterRoleBindingRouter.POST("getClusterRoleBindingDetail", clusterRoleBindingApi.GetClusterRoleBindingDetail)
		clusterRoleBindingRouter.POST("deleteClusterRoleBinding", clusterRoleBindingApi.DeleteClusterRoleBinding)
	}
	return clusterRoleBindingRouter
}
