package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
	k8sResource "github.com/pddzl/kubefish/server/model/k8s/resource"
)

type ResourceApi struct{}

func (ra *ResourceApi) GetResourceRaw(c *gin.Context) {
	var resource k8sRequest.ResourceRaw

	_ = c.ShouldBindJSON(&resource)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&resource); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if raw, err := resourceService.GetResourceRaw(resource.Name, resource.Resource, resource.Namespace); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

// CreateDynamicResource 通过编排动态创建资源
func (ra *ResourceApi) CreateDynamicResource(c *gin.Context) {
	var dynamicResource k8sResource.DynamicResource
	_ = c.ShouldBindJSON(&dynamicResource)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&dynamicResource); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := resourceService.CreateDynamicResource(dynamicResource.Content); err != nil {
		global.KF_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
