package k8s

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
	case "pods":
		req, err = global.KF_K8S_Client.RESTClient().Get().AbsPath("/api/v1").Resource("pods").Name(name).Namespace(namespace).DoRaw(context.TODO())
	}

	return json.RawMessage(req), nil
}
