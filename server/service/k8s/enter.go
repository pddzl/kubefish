package k8s

import (
	"github.com/pddzl/kubefish/server/service/k8s/accessControl"
	"github.com/pddzl/kubefish/server/service/k8s/cluster"
	"github.com/pddzl/kubefish/server/service/k8s/config"
	"github.com/pddzl/kubefish/server/service/k8s/service"
	"github.com/pddzl/kubefish/server/service/k8s/workloads"
)

type ServiceGroup struct {
	ConfigService        config.ServiceGroup
	AccessControlService accessControl.ServiceGroup
	WorkloadsService     workloads.ServiceGroup
	ServiceService       service.ServiceGroup
	ClusterService       cluster.ServiceGroup
}
