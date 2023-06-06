package k8s

import (
	coreV1 "k8s.io/api/core/v1"
	"time"
)

// PodBrief pod列表
type PodBrief struct {
	Name              string    `json:"name,omitempty"`
	Namespace         string    `json:"namespace,omitempty"`
	Status            string    `json:"status,omitempty"`
	Node              string    `json:"node,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}

// PodDetail pod详情
type PodDetail struct {
	MetaData ObjectMeta `json:"metadata"`
	Spec     spec       `json:"spec,omitempty"`
	Status   status     `json:"status,omitempty"`
}

type spec struct {
	InitContainers     []Container `json:"initContainers,omitempty"`
	Containers         []Container `json:"containers"`
	RestartPolicy      string      `json:"restartPolicy,omitempty"`
	ServiceAccountName string      `json:"serviceAccountName,omitempty"`
	NodeName           string      `json:"nodeName,omitempty"`
}

type Container struct {
	Name            string                  `json:"name"`
	Image           string                  `json:"image,omitempty"`
	Command         []string                `json:"command,omitempty"`
	Args            []string                `json:"args,omitempty"`
	VolumeMounts    []VolumeMount           `json:"volumeMounts,omitempty"`
	LivenessProbe   *coreV1.Probe           `json:"livenessProbe,omitempty"`
	ReadinessProbe  *coreV1.Probe           `json:"readinessProbe,omitempty"`
	StartupProbe    *coreV1.Probe           `json:"startupProbe,omitempty"`
	SecurityContext *coreV1.SecurityContext `json:"securityContext,omitempty"`
	Status          *coreV1.ContainerStatus `json:"status,omitempty"`
}

type VolumeMount struct {
	Name       string `json:"name"`
	ReadOnly   bool   `json:"readOnly,omitempty"`
	MountPath  string `json:"mountPath"`
	VolumeType string `json:"volumeType"`
}

type status struct {
	Phase      string                `json:"phase,omitempty"`
	Conditions []coreV1.PodCondition `json:"conditions,omitempty"`
	PodIP      string                `json:"podIP,omitempty"`
	QOSClass   string                `json:"qosClass,omitempty"`
}
