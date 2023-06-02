package request

import "github.com/pddzl/kubefish/server/model/common/request"

// RsListReq 获取replicaSet列表
type RsListReq struct {
	request.PageInfo
	Namespace string `json:"namespace"`
}

type ReplicaSetCommon struct {
	NameSpace  string `json:"namespace" validate:"required"`
	ReplicaSet string `json:"replicaSet" validate:"required"`
}
