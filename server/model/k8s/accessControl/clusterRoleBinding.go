package accessControl

import (
	"time"
)

type ClusterRoleBindingBrief struct {
	Name              string    `json:"name,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}
