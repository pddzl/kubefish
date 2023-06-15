package k8s

import "time"

type IngressBrief struct {
	Name              string    `json:"name"`
	NameSpace         string    `json:"namespace"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	Ip                []string  `json:"ip"`
	Hosts             []string  `json:"hosts"`
}
