package k8s

import "github.com/gin-gonic/gin"

type IngressApi struct{}

// GetIngressList 获取所有ingress
func (ia *IngressApi) GetIngressList(c *gin.Context) {}

// GetIngressDetail 获取ingress 详情
func (ia *IngressApi) GetIngressDetail(c *gin.Context) {}

// DeleteIngress 删除Ingress
func (ia *IngressApi) DeleteIngress(c *gin.Context) {}
