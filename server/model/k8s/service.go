package k8s

import "time"

type ServiceBrief struct {
	Name              string    `json:"name,omitempty"`
	NameSpace         string    `json:"namespace,omitempty"`
	ClusterIP         string    `json:"clusterIP,omitempty"`
	Type              string    `json:"type,omitempty"`
	External          string    `json:"external,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}
