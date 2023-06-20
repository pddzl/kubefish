package config

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	ConfigMapApi
}

var (
	configMapService = service.ServiceGroupApp.K8sServiceGroup.ConfigService.ConfigMapService
)
