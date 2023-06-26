package accessControl

import (
	"time"
)

type ClusterRoleBrief struct {
	Name              string    `json:"name,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}
