package request

import "github.com/pddzl/kubefish/server/model/common/request"

type CommonListReq struct {
	Namespace string `json:"namespace"`
	Label     string `json:"label"`
	Field     string `json:"field"`
	request.PageInfo
}

type CommonReq struct {
	Namespace string `json:"namespace" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type CommonRelatedReq struct {
	Namespace string `json:"namespace" validate:"required"`
	Name      string `json:"name" validate:"required"`
	request.PageInfo
}
