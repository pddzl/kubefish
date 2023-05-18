package router

import (
	"github.com/pddzl/kubefish/server/router/k8s"
	"github.com/pddzl/kubefish/server/router/system"
)

type RouterGroup struct {
	SystemRouterGroup system.RouterGroup
	K8sRouterGroup    k8s.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
