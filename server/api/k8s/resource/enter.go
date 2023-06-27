package resource

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	ResourceApi
}

var (
	resourceService = service.ServiceGroupApp.K8sServiceGroup.ResourceService.ResourceService
)
