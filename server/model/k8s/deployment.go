package k8s

import (
	"k8s.io/apimachinery/pkg/util/intstr"
	"time"
)

type DeploymentBrief struct {
	Name              string    `json:"name"`
	Namespace         string    `json:"namespace"`
	AvailableReplicas int32     `json:"availableReplicas,omitempty"`
	Replicas          int32     `json:"replicas,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}

type DeploymentDetail struct {
	ObjectMeta ObjectMeta       `json:"metadata"`
	Spec       dSpec            `json:"spec"`
	Status     deploymentStatus `json:"status,omitempty"`
}

type dSpec struct {
	Replicas             *int32             `json:"replicas,omitempty"`
	Selector             map[string]string  `json:"selector"`
	Strategy             deploymentStrategy `json:"strategy,omitempty"`
	MinReadySeconds      int32              `json:"minReadySeconds,omitempty"`
	RevisionHistoryLimit *int32             `json:"revisionHistoryLimit,omitempty"`
}

type deploymentStrategy struct {
	DeploymentStrategyType  string                  `json:"type,omitempty"`
	RollingUpdateDeployment rollingUpdateDeployment `json:"rollingUpdate,omitempty"`
}

type rollingUpdateDeployment struct {
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
	MaxSurge       *intstr.IntOrString `json:"maxSurge,omitempty"`
}

type deploymentStatus struct {
	Replicas            int32                 `json:"replicas,omitempty"`
	UpdatedReplicas     int32                 `json:"updatedReplicas,omitempty"`
	ReadyReplicas       int32                 `json:"readyReplicas,omitempty"`
	AvailableReplicas   int32                 `json:"availableReplicas,omitempty"`
	UnavailableReplicas int32                 `json:"unavailableReplicas,omitempty"`
	Conditions          []DeploymentCondition `json:"conditions,omitempty"`
}

type DeploymentCondition struct {
	Type               string    `json:"type"`
	Status             string    `json:"status"`
	LastUpdateTime     time.Time `json:"lastUpdateTime,omitempty"`
	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`
	Reason             string    `json:"reason,omitempty"`
	Message            string    `json:"message,omitempty"`
}
