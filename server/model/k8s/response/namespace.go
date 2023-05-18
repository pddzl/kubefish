package response

import (
	"github.com/pddzl/kubefish/server/model/k8s"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceBrief struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	CreationTimestamp metaV1.Time       `json:"creationTimestamp"`
	Status            string            `json:"status"`
}

type NamespaceDetail struct {
	Metadata       k8s.ObjectMeta  `json:"metadata"`
	Status         string          `json:"status"`
	ResourceQuotas []ResourceQuota `json:"resourceQuotas"`
	LimitRanges    []LimitRange    `json:"limitRanges"`
}

type ResourceQuota struct {
	Name string              `json:"name"`
	Hard coreV1.ResourceList `json:"hard"`
	Used coreV1.ResourceList `json:"used"`
}

type LimitRange struct {
	Name   string                  `json:"name"`
	Limits []coreV1.LimitRangeItem `json:"limits"`
}
