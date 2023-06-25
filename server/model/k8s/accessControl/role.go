package accessControl

import "time"

type RoleBrief struct {
	Name              string    `json:"name,omitempty"`
	Namespace         string    `json:"namespace,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}
