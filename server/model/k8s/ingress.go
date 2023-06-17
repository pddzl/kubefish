package k8s

import "time"

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
	IngressClassName *string       `json:"ingressClassName,omitempty"`
	Rules            []IngressRule `json:"rules,omitempty"`
}

type IngressRule struct {
	Host              string `json:"host"`
	Path              string `json:"path"`
	PathType          string `json:"pathType"`
	ServiceName       string `json:"serviceName"`
	ServicePortName   string `json:"servicePortName"`
	ServicePortNumber int32  `json:"servicePortNumber"`
}

type ingressStatus struct {
	EndPoints []string `json:"endPoints"`
}
