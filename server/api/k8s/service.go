package k8s

import "github.com/gin-gonic/gin"

type ServiceApi struct{}

func (sa *ServiceApi) GetServices(c *gin.Context) {}

func (sa *ServiceApi) GetServiceDetail(c *gin.Context) {}

func (sa *ServiceApi) GetServicePods(c *gin.Context) {}

func (sa *ServiceApi) DeleteService(c *gin.Context) {}
