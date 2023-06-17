package k8s

import (
	"context"
	networkingV1 "k8s.io/api/networking/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8s "github.com/pddzl/kubefish/server/model/k8s"
)

type IngressService struct{}

// GetIngressList 获取所有ingress
func (is *IngressService) GetIngressList(namespace string, pageInfo request.PageInfo) ([]modelK8s.IngressBrief, int, error) {
	// 获取ingress list
	list, err := global.KF_K8S_Client.NetworkingV1().Ingresses(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var ingressList networkingV1.IngressList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		ingressList.Items = list.Items[offset:]
	} else {
		ingressList.Items = list.Items[offset:end]
	}

	// 处理list数据
	var IngressBriefList []modelK8s.IngressBrief
	for _, ingress := range ingressList.Items {
		var ingressBrief modelK8s.IngressBrief
		ingressBrief.Name = ingress.Name
		ingressBrief.NameSpace = ingress.Namespace
		ingressBrief.CreationTimestamp = ingress.CreationTimestamp.Time
		// endPoint
		for _, ig := range ingress.Status.LoadBalancer.Ingress {
			if ig.Hostname != "" {
				ingressBrief.Ip = append(ingressBrief.Ip, ig.Hostname)
			} else if ig.IP != "" {
				ingressBrief.Ip = append(ingressBrief.Ip, ig.IP)
			}
		}
		// hosts
		set := make(map[string]struct{})
		hosts := make([]string, 0)
		for _, rule := range ingress.Spec.Rules {
			if _, exists := set[rule.Host]; !exists && len(rule.Host) > 0 {
				hosts = append(hosts, rule.Host)
			}
			set[rule.Host] = struct{}{}
		}
		ingressBrief.Hosts = hosts
		// append
		IngressBriefList = append(IngressBriefList, ingressBrief)
	}

	return IngressBriefList, total, nil
}

// GetIngressDetail 获取ingress 详情
func (is *IngressService) GetIngressDetail(namespace string, name string) (*modelK8s.IngressDetail, error) {
	detail, err := global.KF_K8S_Client.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理ingress数据
	var ingressDetail modelK8s.IngressDetail
	// metadata
	ingressDetail.ObjectMeta = modelK8s.NewObjectMeta(detail.ObjectMeta)
	// spec
	ingressDetail.Spec.IngressClassName = detail.Spec.IngressClassName
	for _, rule := range detail.Spec.Rules {
		var ingressRule modelK8s.IngressRule
		ingressRule.Host = rule.Host
		for _, path := range rule.IngressRuleValue.HTTP.Paths {
			ingressRule.Path = path.Path
			ingressRule.PathType = string(*path.PathType)
			ingressRule.ServiceName = path.Backend.Service.Name
			ingressRule.ServicePortName = path.Backend.Service.Port.Name
			ingressRule.ServicePortNumber = path.Backend.Service.Port.Number
		}
		ingressDetail.Spec.Rules = append(ingressDetail.Spec.Rules, ingressRule)
	}
	// status
	for _, ig := range detail.Status.LoadBalancer.Ingress {
		if ig.Hostname != "" {
			ingressDetail.Status.EndPoints = append(ingressDetail.Status.EndPoints, ig.Hostname)
		} else if ig.IP != "" {
			ingressDetail.Status.EndPoints = append(ingressDetail.Status.EndPoints, ig.IP)
		}
	}

	return &ingressDetail, nil
}

// DeleteIngress 删除Ingress
func (is *IngressService) DeleteIngress() {}
