package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type ReplicaSetRouter struct{}

func (rs *ReplicaSetRouter) InitReplicaSetRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	replicaSetRouter := Router.Group("replicaSet").Use()
	ReplicaSetApi := api.ApiGroupApp.K8sApiGroup.ReplicaSetApi
	{
		replicaSetRouter.POST("getReplicaSets", ReplicaSetApi.GetReplicaSets)               // 获取所有replicaSet
		replicaSetRouter.POST("getReplicaSetDetail", ReplicaSetApi.GetReplicaSetDetail)     // 获取replicaSet详情
		replicaSetRouter.POST("getReplicaSetPods", ReplicaSetApi.GetReplicaSetPods)         // 获取replicaSet pods
		replicaSetRouter.POST("getReplicaSetServices", ReplicaSetApi.GetReplicaSetServices) // 获取replicaSet services
		replicaSetRouter.POST("deleteReplicaSet", ReplicaSetApi.DeleteReplicaSet)           // 删除replicaSet
	}
	return replicaSetRouter
}
