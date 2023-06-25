package accessControl

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
	"go.uber.org/zap"
)

type RoleBindingApi struct{}

// GetRoleBindingList 获取所有roleBinding
func (rba *RoleBindingApi) GetRoleBindingList(c *gin.Context) {
	var cListReq k8sRequest.CommonListReq
	_ = c.ShouldBindJSON(&cListReq)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&cListReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := roleBindingService.GetRoleBindingList(cListReq.Namespace, cListReq.PageInfo)
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

// GetRoleBindingDetail 获取roleBinding详情
func (rba *RoleBindingApi) GetRoleBindingDetail(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := roleBindingService.GetRoleBindingDetail(commonReq.Namespace, commonReq.Name)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// DeleteRoleBinding 删除roleBinding
func (rba *RoleBindingApi) DeleteRoleBinding(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := roleBindingService.DeleteRoleBinding(commonReq.Namespace, commonReq.Name); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KF_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
