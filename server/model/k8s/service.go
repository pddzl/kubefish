package k8s

import (
	coreV1 "k8s.io/api/core/v1"
	"time"
)

// ServiceBrief list
type ServiceBrief struct {
	Name              string    `json:"name,omitempty"`
	NameSpace         string    `json:"namespace,omitempty"`
	ClusterIP         string    `json:"clusterIP,omitempty"`
	Type              string    `json:"type,omitempty"`
	External          string    `json:"external,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}

// ServiceDetail brief
type ServiceDetail struct {
	ObjectMeta ObjectMeta  `json:"metadata"`
	Spec       serviceSpec `json:"spec,omitempty"`
}

type serviceSpec struct {
	Ports    []coreV1.ServicePort `json:"ports"`
	Selector map[string]string    `json:"selector"`
}
