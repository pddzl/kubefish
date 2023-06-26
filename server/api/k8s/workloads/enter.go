package workloads

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	PodApi
	ReplicaSetApi
	DeploymentApi
	DaemonSetApi
}

var (
	podService        = service.ServiceGroupApp.K8sServiceGroup.WorkloadsService.PodService
	replicaSetService = service.ServiceGroupApp.K8sServiceGroup.WorkloadsService.ReplicaSetService
	deploymentService = service.ServiceGroupApp.K8sServiceGroup.WorkloadsService.DeploymentService
	daemonSetService  = service.ServiceGroupApp.K8sServiceGroup.WorkloadsService.DaemonSetService
)
