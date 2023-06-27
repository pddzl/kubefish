package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
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
