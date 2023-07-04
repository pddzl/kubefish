package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"strings"

	"github.com/pddzl/kubefish/server/global"
)

type ResourceService struct{}

// GetResourceRaw 获取编排内容
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

// CreateDynamicResource 通过编排创建资源
func (rs *ResourceService) CreateDynamicResource(content string) error {
	reader := strings.NewReader(content)
	d := yaml.NewYAMLOrJSONDecoder(reader, 4096)

	data := &unstructured.Unstructured{}
	if err := d.Decode(data); err != nil {
		if err == io.EOF {
			return fmt.Errorf("io EOF")
		}
		return err
	}

	version := data.GetAPIVersion()
	kind := data.GetKind()

	gv, err := schema.ParseGroupVersion(version)
	if err != nil {
		gv = schema.GroupVersion{Version: version}
	}

	//cfg := ctrl.GetConfigOrDie()
	// 创建discovery客户端
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(global.KF_K8S_CONFIG)
	if err != nil {
		return err
	}

	apiResourceList, err := discoveryClient.ServerResourcesForGroupVersion(version)
	if err != nil {
		return err
	}
	apiResources := apiResourceList.APIResources
	var resource *metaV1.APIResource
	for _, apiResource := range apiResources {
		if apiResource.Kind == kind && !strings.Contains(apiResource.Name, "/") {
			resource = &apiResource
			break
		}
	}
	if resource == nil {
		return fmt.Errorf("unknown resource kind: %s", kind)
	}

	// 创建动态客户端
	dynamicClient, err := dynamic.NewForConfig(global.KF_K8S_CONFIG)
	if err != nil {
		return err
	}

	groupVersionResource := schema.GroupVersionResource{Group: gv.Group, Version: gv.Version, Resource: resource.Name}

	namespace := data.GetNamespace()
	if resource.Namespaced {
		_, err = dynamicClient.Resource(groupVersionResource).Namespace(namespace).Create(context.TODO(), data, metaV1.CreateOptions{})
	} else {
		_, err = dynamicClient.Resource(groupVersionResource).Create(context.TODO(), data, metaV1.CreateOptions{})
	}

	return err
}
