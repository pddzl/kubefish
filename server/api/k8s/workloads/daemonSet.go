package workloads

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
	"go.uber.org/zap"
)

type DaemonSetApi struct{}

// GetDaemonSets 获取所有daemonSet
func (da *DaemonSetApi) GetDaemonSets(c *gin.Context) {
	var cListReq k8sRequest.CommonListReq
	_ = c.ShouldBindJSON(&cListReq)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&cListReq); err != nil {
		global.KF_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := daemonSetService.GetDaemonSets(cListReq.Namespace, cListReq.Label, cListReq.Field, cListReq.PageInfo)
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

// GetDaemonSetDetail 获取daemonSet 详情
func (da *DaemonSetApi) GetDaemonSetDetail(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := daemonSetService.GetDaemonSetDetail(commonReq.Namespace, commonReq.Name)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// GetDaemonSetPods 获取daemonSet关联的pod
func (da *DaemonSetApi) GetDaemonSetPods(c *gin.Context) {
	var commonRelatedReq k8sRequest.CommonRelatedReq
	_ = c.ShouldBindJSON(&commonRelatedReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonRelatedReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := daemonSetService.GetDaemonSetPods(commonRelatedReq.Namespace, commonRelatedReq.Name, commonRelatedReq.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     commonRelatedReq.Page,
			PageSize: commonRelatedReq.PageSize,
		}, "获取成功", c)
	}
}

// DeleteDaemonSet 删除daemonSet
func (da *DaemonSetApi) DeleteDaemonSet(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := daemonSetService.DeleteDaemonSet(commonReq.Namespace, commonReq.Name); err != nil {
		response.FailWithMessage("删除失败", c)
		global.KF_LOG.Error("删除失败", zap.Error(err))
		return
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
