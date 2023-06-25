package accessControl

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	ServiceAccountApi
	RoleApi
}

var (
	serviceAccountService = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.ServiceAccountService
	roleService           = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.RoleService
)
