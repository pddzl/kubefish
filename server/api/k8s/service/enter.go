package service

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	ServicesApi
	IngressApi
}

var (
	servicesService = service.ServiceGroupApp.K8sServiceGroup.ServiceService.ServicesService
	ingressService  = service.ServiceGroupApp.K8sServiceGroup.ServiceService.IngressService
)
