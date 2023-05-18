package system

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type CasbinRouter struct{}

func (cr *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	casbinRouter := Router.Group("casbin")
	casbinApi := api.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.POST("editCasbin", casbinApi.EditCasbin)
	}
	return casbinRouter
}
