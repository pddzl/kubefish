package config

import (
	"time"
)

type SecretBrief struct {
	Name              string    `json:"name"`
	Namespace         string    `json:"namespace"`
	Type              string    `json:"type"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}
