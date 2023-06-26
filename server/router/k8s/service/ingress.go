package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type IngressRouter struct{}

func (ir *IngressRouter) InitIngressRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	ingressRouter := Router.Group("ingress")
	ingressApi := api.ApiGroupApp.K8sApiGroup.ServiceApi.IngressApi
	{
		ingressRouter.POST("getIngressList", ingressApi.GetIngressList)     // 获取所有ingress
		ingressRouter.POST("getIngressDetail", ingressApi.GetIngressDetail) // 获取ingress详情
		ingressRouter.POST("deleteIngress", ingressApi.DeleteIngress)       // 删除ingress
	}
	return ingressRouter
}
