package accessControl

type ServiceGroup struct {
	ServiceAccountService
	RoleService
	RoleBindingService
	ClusterRoleService
	ClusterRoleBindingService
}
