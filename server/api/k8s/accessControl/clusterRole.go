package accessControl

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	"github.com/pddzl/kubefish/server/model/common/response"
)

type ClusterRoleApi struct{}

// GetClusterRoleList 获取所有clusterRole
func (cra *ClusterRoleApi) GetClusterRoleList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := clusterRoleService.GetClusterRoleList(pageInfo)
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

// GetClusterRoleDetail 获取clusterRole详情
func (cra *ClusterRoleApi) GetClusterRoleDetail(c *gin.Context) {
	var clusterRole request.GetByName
	_ = c.ShouldBindJSON(&clusterRole)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRole); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := clusterRoleService.GetClusterRoleDetail(clusterRole.Name)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

// DeleteClusterRole 删除clusterRole
func (cra *ClusterRoleApi) DeleteClusterRole(c *gin.Context) {
	var clusterRole request.GetByName
	_ = c.ShouldBindJSON(&clusterRole)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRole); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clusterRoleService.DeleteClusterRole(clusterRole.Name); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KF_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
