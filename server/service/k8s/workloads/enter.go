package workloads

type ServiceGroup struct {
	PodService
	ReplicaSetService
	DeploymentService
	DaemonSetService
}
