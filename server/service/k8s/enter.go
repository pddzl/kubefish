package k8s

type ServiceGroup struct {
	NodeService
	ResourceService
	NamespaceService
	PodService
	ReplicaSetService
	DeploymentService
	DaemonSetService
}
