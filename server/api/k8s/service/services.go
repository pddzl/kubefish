package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
)

type ServicesApi struct{}

// GetServices 获取所有service
func (sa *ServicesApi) GetServices(c *gin.Context) {
	var cListReq k8sRequest.CommonListReq
	_ = c.ShouldBindJSON(&cListReq)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&cListReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := servicesService.GetServices(cListReq.Namespace, cListReq.Label, cListReq.Field, cListReq.PageInfo)
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

// GetServiceDetail 获取service详情
func (sa *ServicesApi) GetServiceDetail(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := servicesService.GetServiceDetail(commonReq.Namespace, commonReq.Name)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

func (sa *ServicesApi) GetServicePods(c *gin.Context) {
	var relatedReq k8sRequest.CommonRelatedReq
	_ = c.ShouldBindJSON(&relatedReq)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&relatedReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pods, total, err := servicesService.GetServicePods(relatedReq.Namespace, relatedReq.Name, relatedReq.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     pods,
		Total:    int64(total),
		Page:     relatedReq.Page,
		PageSize: relatedReq.PageSize,
	}, "获取成功", c)
}

// DeleteService 删除服务
func (sa *ServicesApi) DeleteService(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := servicesService.DeleteService(commonReq.Namespace, commonReq.Name); err != nil {
		response.FailWithMessage("删除失败", c)
		global.KF_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
