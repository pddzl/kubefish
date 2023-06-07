package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type DeploymentRouter struct{}

func (dr *DeploymentRouter) InitDeploymentRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	deploymentRouter := Router.Group("deployment")
	deploymentApi := api.ApiGroupApp.K8sApiGroup.DeploymentApi
	{
		deploymentRouter.POST("getDeployments", deploymentApi.GetDeployments)           // 获取所有deployment
		deploymentRouter.POST("getDeploymentDetail", deploymentApi.GetDeploymentDetail) // 获取deployment详情
		deploymentRouter.POST("getDeploymentRs", deploymentApi.GetDeploymentRs)         // 获取deployment rs
		deploymentRouter.POST("deleteDeployment", deploymentApi.DeleteDeployment)       // 删除deployment
		deploymentRouter.POST("scaleDeployment", deploymentApi.ScaleDeployment)         // 伸缩
	}
	return deploymentRouter
}
