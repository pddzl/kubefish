package accessControl

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	ServiceAccountApi
}

var (
	serviceAccountService = service.ServiceGroupApp.K8sServiceGroup.AccessControlService.ServiceAccountService
)
