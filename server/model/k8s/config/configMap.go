package config

import (
	"time"
)

type ConfigMapBrief struct {
	Name              string    `json:"name"`
	Namespace         string    `json:"namespace"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}
