package accessControl

type RouterGroup struct {
	ServiceAccountRouter
	RoleRouter
	RoleBindingRouter
	ClusterRoleRouter
	ClusterRoleBindingRouter
}
