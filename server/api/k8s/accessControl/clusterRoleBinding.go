package accessControl

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	"github.com/pddzl/kubefish/server/model/common/response"
)

type ClusterRoleBindingApi struct{}

// GetClusterRoleBindingList 获取所有clusterRoleBinding
func (cb *ClusterRoleBindingApi) GetClusterRoleBindingList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := clusterRoleBindingService.GetClusterRoleBindingList(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetClusterRoleBindingDetail 获取clusterRoleBinding详情
func (cb *ClusterRoleBindingApi) GetClusterRoleBindingDetail(c *gin.Context) {
	var clusterRoleBinding request.GetByName
	_ = c.ShouldBindJSON(&clusterRoleBinding)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRoleBinding); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := clusterRoleBindingService.GetClusterRoleBinding(clusterRoleBinding.Name)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

// DeleteClusterRoleBinding 删除ClusterRoleBinding
func (cb *ClusterRoleBindingApi) DeleteClusterRoleBinding(c *gin.Context) {
	var clusterRoleBinding request.GetByName
	_ = c.ShouldBindJSON(&clusterRoleBinding)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRoleBinding); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clusterRoleBindingService.DeleteClusterRoleBinding(clusterRoleBinding.Name); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KF_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
