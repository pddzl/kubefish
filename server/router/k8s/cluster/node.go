package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type NodeRouter struct{}

func (nr *NodeRouter) InitNodeRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	nodeRouter := Router.Group("node")
	nodeApi := api.ApiGroupApp.K8sApiGroup.ClusterApi.NodeApi
	{
		nodeRouter.POST("getNodes", nodeApi.GetNodes)
		nodeRouter.GET("getNodeDetail", nodeApi.GetNodeDetail)
		nodeRouter.POST("getNodePods", nodeApi.GetNodePods)
	}
	return nodeRouter
}
