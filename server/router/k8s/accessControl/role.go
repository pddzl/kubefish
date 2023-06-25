package accessControl

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type RoleRouter struct{}

func (rr *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	roleRouter := Router.Group("role")
	roleApi := api.ApiGroupApp.K8sApiGroup.AccessControlApi.RoleApi
	{
		roleRouter.POST("getRoleList", roleApi.GetRoleList)
		roleRouter.POST("getRoleDetail", roleApi.GetRoleDetail)
		roleRouter.POST("deleteRole", roleApi.DeleteRole)
	}
	return roleRouter
}
