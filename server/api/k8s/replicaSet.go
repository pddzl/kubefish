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
	var commonList k8sRequest.CommonListReq
	_ = c.ShouldBindJSON(&commonList)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&commonList); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := replicaSetService.GetReplicaSets(commonList.Namespace, commonList.Label, commonList.Field, commonList.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(total),
		Page:     commonList.Page,
		PageSize: commonList.PageSize,
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

	detail, err := replicaSetService.GetReplicaSetDetail(replicaSetCommon.Namespace, replicaSetCommon.ReplicaSet)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// GetReplicaSetPods 获取replicaSet关联的pods
func (rs *ReplicaSetApi) GetReplicaSetPods(c *gin.Context) {
	var rsPods k8sRequest.RsPods
	_ = c.ShouldBindJSON(&rsPods)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&rsPods); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pods, total, err := replicaSetService.GetReplicaSetPods(rsPods.Namespace, rsPods.ReplicaSet, rsPods.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     pods,
		Total:    int64(total),
		Page:     rsPods.Page,
		PageSize: rsPods.PageSize,
	}, "获取成功", c)
}

// DeleteReplicaSet 删除replicaSet
func (rs *ReplicaSetApi) DeleteReplicaSet(c *gin.Context) {
	var replicaSet k8sRequest.ReplicaSetCommon
	_ = c.ShouldBindJSON(&replicaSet)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&replicaSet); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := replicaSetService.DeleteReplicaSet(replicaSet.Namespace, replicaSet.ReplicaSet); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KF_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
