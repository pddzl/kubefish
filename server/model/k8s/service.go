package k8s

import (
	coreV1 "k8s.io/api/core/v1"
	"time"
)

// ServiceBrief list
type ServiceBrief struct {
	Name              string    `json:"name"`
	NameSpace         string    `json:"namespace"`
	ClusterIP         string    `json:"clusterIP"`
	Type              string    `json:"type"`
	External          string    `json:"external"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
}

// ServiceDetail brief
type ServiceDetail struct {
	ObjectMeta ObjectMeta  `json:"metadata"`
	Spec       serviceSpec `json:"spec"`
}

type serviceSpec struct {
	Ports           []coreV1.ServicePort `json:"ports"`
	Type            string               `json:"type"`
	ClusterIP       string               `json:"clusterIP"`
	SessionAffinity string               `json:"sessionAffinity"`
	Selector        map[string]string    `json:"selector"`
}
