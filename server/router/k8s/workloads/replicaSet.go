package workloads

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ReplicaSetRouter struct{}

func (rs *ReplicaSetRouter) InitReplicaSetRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	replicaSetRouter := Router.Group("replicaSet")
	ReplicaSetApi := api.ApiGroupApp.K8sApiGroup.WorkloadsApi.ReplicaSetApi
	{
		replicaSetRouter.POST("getReplicaSets", ReplicaSetApi.GetReplicaSets)           // 获取所有replicaSet
		replicaSetRouter.POST("getReplicaSetDetail", ReplicaSetApi.GetReplicaSetDetail) // 获取replicaSet详情
		replicaSetRouter.POST("getReplicaSetPods", ReplicaSetApi.GetReplicaSetPods)     // 获取replicaSet pods
		replicaSetRouter.POST("deleteReplicaSet", ReplicaSetApi.DeleteReplicaSet)       // 删除replicaSet
	}
	return replicaSetRouter
}
