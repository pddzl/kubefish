package k8s

type RouterGroup struct {
	NodeRouter
	ResourceRouter
	NamespaceRouter
	PodRouter
	ReplicaSetRouter
	DeploymentRouter
	DaemonSetRouter
	ServiceRouter
}
