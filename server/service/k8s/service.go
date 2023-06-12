package k8s

import (
	"context"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8s "github.com/pddzl/kubefish/server/model/k8s"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceService struct {
}

// GetServices 获取所有service
func (ss *ServiceService) GetServices(namespace string, label string, field string, pageInfo request.PageInfo) ([]modelK8s.ServiceBrief, int, error) {
	// 获取service list
	var option metaV1.ListOptions
	if label != "" {
		option.LabelSelector = label
	}
	if field != "" {
		option.FieldSelector = field
	}

	list, err := global.KF_K8S_Client.CoreV1().Services(namespace).List(context.TODO(), option)
	if err != nil {
		return nil, 0, err
	}

	var serviceList coreV1.ServiceList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		serviceList.Items = list.Items[offset:]
	} else {
		serviceList.Items = list.Items[offset:end]
	}

	var serviceBriefList []modelK8s.ServiceBrief
	// 处理list数据
	for _, slr := range serviceList.Items {
		var servicesBrief modelK8s.ServiceBrief
		servicesBrief.Name = slr.Name
		servicesBrief.NameSpace = slr.Namespace
		servicesBrief.ClusterIP = slr.Spec.ClusterIP
		servicesBrief.Type = string(slr.Spec.Type)
		servicesBrief.CreationTimestamp = slr.CreationTimestamp.Time
		// append
		serviceBriefList = append(serviceBriefList, servicesBrief)
	}

	return serviceBriefList, total, nil
}

func (ss *ServiceService) GetServiceDetail() {}

func (ss *ServiceService) GetServicePods() {}

// DeleteService 删除服务
func (ss *ServiceService) DeleteService(namespace string, name string) error {
	return global.KF_K8S_Client.CoreV1().Services(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
