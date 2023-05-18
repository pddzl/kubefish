package service

import (
	"github.com/pddzl/kubefish/server/service/k8s"
	"github.com/pddzl/kubefish/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	K8sServiceGroup    k8s.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
