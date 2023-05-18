package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	"github.com/pddzl/kubefish/server/model/common/response"
	k8sRequest "github.com/pddzl/kubefish/server/model/k8s/request"
	"go.uber.org/zap"
)

type NodeApi struct{}

// GetNodes 获取节点列表
func (na *NodeApi) GetNodes(c *gin.Context) {
	var info request.PageInfo

	_ = c.ShouldBindJSON(&info)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := nodeService.GetNodes(info); err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     info.Page,
			PageSize: info.PageSize,
		}, "获取成功", c)
	}
}

// GetNodeDetail 获取某个节点详情
func (na *NodeApi) GetNodeDetail(c *gin.Context) {
	name := c.Query("name")

	if detail, err := nodeService.GetNodeDetail(name); err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

func (na *NodeApi) GetNodePods(c *gin.Context) {
	var node k8sRequest.NodePods
	_ = c.ShouldBindJSON(&node)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&node); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pods, total, err := nodeService.GetNodePods(node.NodeName, node.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KF_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     pods,
			Total:    total,
			Page:     node.Page,
			PageSize: node.PageSize,
		}, "获取成功", c)
	}
}
