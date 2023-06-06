package k8s

import "time"

type ReplicaSetBrief struct {
	Name              string    `json:"name"`
	NameSpace         string    `json:"namespace"`
	AvailableReplicas int32     `json:"availableReplicas,omitempty"`
	Replicas          int32     `json:"replicas,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}

type ReplicaSetDetail struct {
	ObjectMeta ObjectMeta       `json:"metadata"`
	Spec       replicaSetSpec   `json:"spec,omitempty"`
	Status     replicaSetStatus `json:"status,omitempty"`
}

type replicaSetSpec struct {
	Replicas        *int32            `json:"replicas,omitempty"`
	MinReadySeconds int32             `json:"minReadySeconds,omitempty"`
	Selector        map[string]string `json:"selector"`
}

type replicaSetStatus struct {
	Replicas             int32                 `json:"replicas"`
	FullyLabeledReplicas int32                 `json:"fullyLabeledReplicas,omitempty"`
	ReadyReplicas        int32                 `json:"readyReplicas,omitempty"`
	AvailableReplicas    int32                 `json:"availableReplicas,omitempty"`
	Conditions           []ReplicaSetCondition `json:"conditions,omitempty"`
}

type ReplicaSetCondition struct {
	Type               string    `json:"type"`
	Status             string    `json:"status"`
	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`
	Reason             string    `json:"reason,omitempty"`
	Message            string    `json:"message,omitempty"`
}

type NewReplicaSet struct {
	Name              string            `json:"name"`
	NameSpace         string            `json:"namespace"`
	Labels            map[string]string `json:"labels,omitempty"`
	CreationTimestamp time.Time         `json:"creationTimestamp,omitempty"`
	Replicas          string            `json:"replicas"`
}
