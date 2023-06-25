package k8s

import (
	"github.com/pddzl/kubefish/server/service/k8s/accessControl"
	"github.com/pddzl/kubefish/server/service/k8s/config"
)

type ServiceGroup struct {
	NodeService
	ResourceService
	NamespaceService
	PodService
	ReplicaSetService
	DeploymentService
	DaemonSetService
	ServiceService
	IngressService
	ConfigService        config.ServiceGroup
	AccessControlService accessControl.ServiceGroup
}
