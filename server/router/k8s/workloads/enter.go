package workloads

type RouterGroup struct {
	PodRouter
	ReplicaSetRouter
	DeploymentRouter
	DaemonSetRouter
}
