package k8s

import (
	"github.com/pddzl/kubefish/server/router/k8s/accessControl"
	"github.com/pddzl/kubefish/server/router/k8s/cluster"
	"github.com/pddzl/kubefish/server/router/k8s/config"
	"github.com/pddzl/kubefish/server/router/k8s/service"
	"github.com/pddzl/kubefish/server/router/k8s/workloads"
)

type RouterGroup struct {
	ConfigRouter        config.RouterGroup
	AccessControlRouter accessControl.RouterGroup
	WorkloadsRouter     workloads.RouterGroup
	ServiceRouter       service.RouterGroup
	ClusterRouter       cluster.RouterGroup
}
