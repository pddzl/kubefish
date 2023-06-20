package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
	"go.uber.org/zap"
)

type ConfigMapApi struct{}

// GetConfigMapList 获取所有configMap
func (ca *ConfigMapApi) GetConfigMapList(c *gin.Context) {
	var cListReq k8sRequest.CommonListReq
	_ = c.ShouldBindJSON(&cListReq)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&cListReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := configMapService.GetConfigMapList(cListReq.Namespace, cListReq.PageInfo)
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

// GetConfigMapDetail 获取单个configMap详情
func (ca *ConfigMapApi) GetConfigMapDetail(c *gin.Context) {}

// DeleteConfigMap 删除configMap
func (ca *ConfigMapApi) DeleteConfigMap(c *gin.Context) {}
