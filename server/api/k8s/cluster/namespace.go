package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	"github.com/pddzl/kubefish/server/model/common/response"
	"go.uber.org/zap"
)

type NamespaceApi struct{}

// GetNamespaces 获取集群所有namespace
func (na *NamespaceApi) GetNamespaces(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := namespaceService.GetNamespaceList(pageInfo)
	if err != nil {
		global.KF_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetNamespaceDetail 获取指定namespace详情
func (na *NamespaceApi) GetNamespaceDetail(c *gin.Context) {
	name := c.Query("name")

	if detail, err := namespaceService.GetNamespaceDetail(name); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

// DeleteNamespace 删除命名空间
func (na *NamespaceApi) DeleteNamespace(c *gin.Context) {
	var byName request.GetByName
	_ = c.ShouldBindJSON(&byName)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&byName); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := namespaceService.DeleteNamespace(byName.Name); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KF_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetNamespaceName 获取集群所有namespace（只包括name）
func (na *NamespaceApi) GetNamespaceName(c *gin.Context) {
	list, err := namespaceService.GetNamespaceName()
	if err != nil {
		global.KF_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}
