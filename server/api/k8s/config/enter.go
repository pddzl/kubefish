package config

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	ConfigMapApi
	SecretApi
}

var (
	configMapService = service.ServiceGroupApp.K8sServiceGroup.ConfigService.ConfigMapService
	secretService    = service.ServiceGroupApp.K8sServiceGroup.ConfigService.SecretService
)
