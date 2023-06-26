package workloads

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
	"go.uber.org/zap"
)

type DeploymentApi struct{}

// GetDeployments 获取所有deployment
func (da *DeploymentApi) GetDeployments(c *gin.Context) {
	var cListReq k8sRequest.CommonListReq
	_ = c.ShouldBindJSON(&cListReq)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&cListReq); err != nil {
		global.KF_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := deploymentService.GetDeploymentList(cListReq.Namespace, cListReq.Label, cListReq.Field, cListReq.PageInfo)
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

// GetDeploymentDetail 获取deployment详情
func (da *DeploymentApi) GetDeploymentDetail(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := deploymentService.GetDeploymentDetail(commonReq.Namespace, commonReq.Name)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// GetDeploymentRs 获取deployment关联的replicaSet
func (da *DeploymentApi) GetDeploymentRs(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if rs, err := deploymentService.GetDeploymentRs(commonReq.Namespace, commonReq.Name); err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(rs, "获取成功", c)
	}
}

// DeleteDeployment 删除deployment
func (da *DeploymentApi) DeleteDeployment(c *gin.Context) {
	var commonReq k8sRequest.CommonReq
	_ = c.ShouldBindJSON(&commonReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&commonReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := deploymentService.DeleteDeployment(commonReq.Namespace, commonReq.Name); err != nil {
		response.FailWithMessage("删除失败", c)
		global.KF_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// ScaleDeployment 伸缩
func (da *DeploymentApi) ScaleDeployment(c *gin.Context) {
	var scale k8sRequest.DmScale
	_ = c.ShouldBindJSON(&scale)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&scale); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := deploymentService.ScaleDeployment(scale.Namespace, scale.Name, scale.Num); err != nil {
		response.FailWithMessage("伸缩失败", c)
		global.KF_LOG.Error("伸缩失败", zap.Error(err))
	} else {
		response.OkWithMessage("伸缩成功", c)
	}
}
