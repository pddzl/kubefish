package k8s

import (
	networkingV1 "k8s.io/api/networking/v1"
	"time"
)

// IngressBrief list
type IngressBrief struct {
	Name              string    `json:"name"`
	NameSpace         string    `json:"namespace"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	Ip                []string  `json:"ip"`
	Hosts             []string  `json:"hosts"`
}

// IngressDetail detail
type IngressDetail struct {
	ObjectMeta ObjectMeta    `json:"metadata"`
	Spec       ingressSpec   `json:"spec"`
	Status     ingressStatus `json:"status,omitempty"`
}

type ingressSpec struct {
	IngressClassName *string                    `json:"ingressClassName,omitempty"`
	Rules            []networkingV1.IngressRule `json:"rules,omitempty"`
}

type ingressStatus struct {
	EndPoints []string `json:"endPoints"`
}
