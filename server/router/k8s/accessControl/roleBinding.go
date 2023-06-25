package accessControl

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type RoleBindingRouter struct{}

func (rbr *RoleBindingRouter) InitRoleBindingRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	roleBindingRouter := Router.Group("roleBinding")
	roleBindingApi := api.ApiGroupApp.K8sApiGroup.AccessControlApi.RoleBindingApi
	{
		roleBindingRouter.POST("getRoleBindingList", roleBindingApi.GetRoleBindingList)
		roleBindingRouter.POST("getRoleBindingDetail", roleBindingApi.GetRoleBindingDetail)
		roleBindingRouter.POST("deleteRoleBinding", roleBindingApi.DeleteRoleBinding)
	}
	return roleBindingRouter
}
