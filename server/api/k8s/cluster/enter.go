package cluster

import (
	"github.com/pddzl/kubefish/server/service"
)

type ApiGroup struct {
	NamespaceApi
	NodeApi
}

var (
	namespaceService = service.ServiceGroupApp.K8sServiceGroup.ClusterService.NamespaceService
	nodeService      = service.ServiceGroupApp.K8sServiceGroup.ClusterService.NodeService
)
