package resource

import (
	"context"
	"encoding/json"
	"github.com/pddzl/kubefish/server/global"
)

type ResourceService struct{}

func (rs *ResourceService) GetResourceRaw(name string, resource string, namespace string) (raw interface{}, err error) {
	var req []byte
	switch resource {
	case "nodes", "namespaces":
		req, err = global.KF_K8S_Client.RESTClient().Get().AbsPath("/api/v1").Resource(resource).Name(name).DoRaw(context.TODO())
	case "pods", "services", "configmaps", "secrets", "serviceaccounts":
		req, err = global.KF_K8S_Client.RESTClient().Get().AbsPath("/api/v1").Resource(resource).Name(name).Namespace(namespace).DoRaw(context.TODO())
	case "replicasets", "deployments", "daemonsets":
		//kubectl get --raw "/apis/apps/v1/namespaces/kube-system/replicasets/coredns-6d8c4cb4d"
		req, err = global.KF_K8S_Client.RESTClient().Get().AbsPath("/apis/apps/v1").Resource(resource).Name(name).Namespace(namespace).DoRaw(context.TODO())
	case "ingresses":
		req, err = global.KF_K8S_Client.RESTClient().Get().AbsPath("/apis/networking.k8s.io/v1").Resource(resource).Name(name).Namespace(namespace).DoRaw(context.TODO())
	case "roles", "rolebindings":
		req, err = global.KF_K8S_Client.RESTClient().Get().AbsPath("/apis/rbac.authorization.k8s.io/v1").Resource(resource).Namespace(namespace).Name(name).DoRaw(context.TODO())
	case "clusterroles", "clusterrolebindings":
		req, err = global.KF_K8S_Client.RESTClient().Get().AbsPath("/apis/rbac.authorization.k8s.io/v1").Resource(resource).Name(name).DoRaw(context.TODO())
	}

	return json.RawMessage(req), nil
}