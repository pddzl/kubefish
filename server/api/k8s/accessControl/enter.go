package accessControl

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	ServiceAccountApi
	RoleApi
	RoleBindingApi
	ClusterRoleApi
	ClusterRoleBindingApi
}

var (
	serviceAccountService     = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.ServiceAccountService
	roleService               = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.RoleService
	roleBindingService        = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.RoleBindingService
	clusterRoleService        = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.ClusterRoleService
	clusterRoleBindingService = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.ClusterRoleBindingService
)
