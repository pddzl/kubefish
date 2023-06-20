package config

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type SecretRouter struct{}

func (sr *SecretRouter) InitSecretRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	secretRouter := Router.Group("secret")
	secretApi := api.ApiGroupApp.K8sApiGroup.ConfigApi.SecretApi
	{
		secretRouter.POST("getSecretList", secretApi.GetSecretList)
		secretRouter.POST("getSecretDetail", secretApi.GetSecretDetail)
		secretRouter.POST("deleteSecret", secretApi.DeleteSecret)
	}
	return secretRouter
}
