package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type DaemonSetRouter struct{}

func (dr *DaemonSetRouter) InitDaemonSetRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	daemonSetRouter := Router.Group("daemonSet")
	daemonSetApi := api.ApiGroupApp.K8sApiGroup.DaemonSetApi
	{
		daemonSetRouter.POST("getDaemonSets", daemonSetApi.GetDaemonSets)           // 获取所有daemonSet
		daemonSetRouter.POST("getDaemonSetDetail", daemonSetApi.GetDaemonSetDetail) // 获取daemonSet详情
		daemonSetRouter.POST("getDaemonSetPods", daemonSetApi.GetDaemonSetPods)     // 获取daemonSet关联pod
		daemonSetRouter.POST("deleteDaemonSet", daemonSetApi.DeleteDaemonSet)       // 删除daemonSet
	}
	return daemonSetRouter
}
