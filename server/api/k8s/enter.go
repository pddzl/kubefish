package k8s

import (
	"github.com/pddzl/kubefish/server/api/k8s/accessControl"
	"github.com/pddzl/kubefish/server/api/k8s/cluster"
	"github.com/pddzl/kubefish/server/api/k8s/config"
	"github.com/pddzl/kubefish/server/api/k8s/service"
	"github.com/pddzl/kubefish/server/api/k8s/workloads"
)

type ApiGroup struct {
	ConfigApi        config.ApiGroup
	AccessControlApi accessControl.ApiGroup
	WorkloadsApi     workloads.ApiGroup
	ServiceApi       service.ApiGroup
	ClusterApi       cluster.ApiGroup
}
