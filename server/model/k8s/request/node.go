package request

import "github.com/pddzl/kubefish/server/model/common/request"

type NodePods struct {
	NodeName string `json:"node_name" validate:"required"`
	request.PageInfo
}
