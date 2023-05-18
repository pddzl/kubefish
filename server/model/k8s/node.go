package k8s

import (
	coreV1 "k8s.io/api/core/v1"
)

type NodeDetail struct {
	Metadata               ObjectMeta             `json:"metadata"`
	Spec                   nodeSpec               `json:"spec,omitempty"`
	Status                 NodeStatus             `json:"status,omitempty"`
	NodeAllocatedResources nodeAllocatedResources `json:"allocatedResources"`
}

type nodeSpec struct {
	PodCIDR       string         `json:"podCIDR,omitempty"`
	Unschedulable bool           `json:"unschedulable,omitempty"`
	Taints        []coreV1.Taint `json:"taints,omitempty"`
}

type NodeStatus struct {
	Phase           string                     `json:"phase,omitempty"`
	Conditions      []coreV1.NodeCondition     `json:"conditions,omitempty"`
	Addresses       []coreV1.NodeAddress       `json:"addresses,omitempty"`
	DaemonEndpoints coreV1.NodeDaemonEndpoints `json:"Port"` // kubelet服务端口
	NodeInfo        coreV1.NodeSystemInfo      `json:"nodeInfo,omitempty"`
	Images          []coreV1.ContainerImage    `json:"images,omitempty"`
}

type nodeAllocatedResources struct {
	CPURequests           int64   `json:"cpuRequests"`
	CPURequestsPercent    float64 `json:"cpuRequestsPercent"`
	CPULimits             int64   `json:"cpuLimits"`
	CPULimitsPercent      float64 `json:"cpuLimitsPercent"`
	AllocatedCPU          int64   `json:"allocatedCPU"`
	MemoryRequests        int64   `json:"memoryRequests"`
	MemoryRequestsPercent float64 `json:"memoryRequestsPercent"`
	MemoryLimits          int64   `json:"memoryLimits"`
	MemoryLimitsPercent   float64 `json:"memoryLimitsPercent"`
	AllocatedMemory       int64   `json:"allocatedMemory"`
	AllocatedPods         int64   `json:"allocatedPods"`
	PodNum                int     `json:"podNum"`
	PodPercent            float64 `json:"podPercent"`
}
