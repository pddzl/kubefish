package accessControl

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	ServiceAccountApi
	RoleApi
	RoleBindingApi
}

var (
	serviceAccountService = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.ServiceAccountService
	roleService           = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.RoleService
	roleBindingService    = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.RoleBindingService
)
