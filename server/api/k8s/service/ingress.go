package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
	"go.uber.org/zap"
)

type IngressApi struct{}

// GetIngressList 获取所有ingress
func (ia *IngressApi) GetIngressList(c *gin.Context) {
	var cListReq k8sRequest.CommonListReq
	_ = c.ShouldBindJSON(&cListReq)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&cListReq); err != nil {
		global.KF_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := ingressService.GetIngressList(cListReq.Namespace, cListReq.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     cListReq.Page,
			PageSize: cListReq.PageSize,
		}, "获取成功", c)
	}
}

// GetIngressDetail 获取ingress 详情
func (ia *IngressApi) GetIngressDetail(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := ingressService.GetIngressDetail(commonReq.Namespace, commonReq.Name)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// DeleteIngress 删除Ingress
func (ia *IngressApi) DeleteIngress(c *gin.Context) {}
