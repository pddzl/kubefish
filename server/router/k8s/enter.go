package k8s

import (
	"github.com/pddzl/kubefish/server/router/k8s/accessControl"
	"github.com/pddzl/kubefish/server/router/k8s/config"
)

type RouterGroup struct {
	NodeRouter
	ResourceRouter
	NamespaceRouter
	PodRouter
	ReplicaSetRouter
	DeploymentRouter
	DaemonSetRouter
	ServiceRouter
	IngressRouter
	ConfigRouter        config.RouterGroup
	AccessControlRouter accessControl.RouterGroup
}
