package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
)

type ReplicaSetApi struct{}

// GetReplicaSets 获取replicaSet列表
func (rs *ReplicaSetApi) GetReplicaSets(c *gin.Context) {
	var pageInfo k8sRequest.RsListReq
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := replicaSetService.GetReplicaSets(pageInfo.Namespace, pageInfo.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(total),
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetReplicaSetDetail 获取replicaSet详情
func (rs *ReplicaSetApi) GetReplicaSetDetail(c *gin.Context) {
	var replicaSetCommon k8sRequest.ReplicaSetCommon
	_ = c.ShouldBindJSON(&replicaSetCommon)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&replicaSetCommon); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := replicaSetService.GetReplicaSetDetail(replicaSetCommon.NameSpace, replicaSetCommon.ReplicaSet)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// GetReplicaSetPods 获取replicaSet关联的pods
func (rs *ReplicaSetApi) GetReplicaSetPods(c *gin.Context) {}

// GetReplicaSetServices 获取replicaSet关联的services
func (rs *ReplicaSetApi) GetReplicaSetServices(c *gin.Context) {}

// DeleteReplicaSet 删除replicaSet
func (rs *ReplicaSetApi) DeleteReplicaSet(c *gin.Context) {}
