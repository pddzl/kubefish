package k8s

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type ObjectMeta struct {
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	CreationTimestamp metaV1.Time       `json:"creationTimestamp,omitempty"`
	UID               types.UID         `json:"uid,omitempty"`
}

func NewObjectMeta(k8SObjectMeta metaV1.ObjectMeta) ObjectMeta {
	return ObjectMeta{
		Name:              k8SObjectMeta.Name,
		Namespace:         k8SObjectMeta.Namespace,
		Labels:            k8SObjectMeta.Labels,
		CreationTimestamp: k8SObjectMeta.CreationTimestamp,
		Annotations:       k8SObjectMeta.Annotations,
		UID:               k8SObjectMeta.UID,
	}
}
