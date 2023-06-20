package config

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8sConfig "github.com/pddzl/kubefish/server/model/k8s/config"
)

type ConfigMapService struct{}

// GetConfigMapList 获取所有configMap
func (ca *ConfigMapService) GetConfigMapList(namespace string, pageInfo request.PageInfo) ([]modelK8sConfig.ConfigMapBrief, int, error) {
	// 获取configMap list
	list, err := global.KF_K8S_Client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var configMapList coreV1.ConfigMapList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		configMapList.Items = list.Items[offset:]
	} else {
		configMapList.Items = list.Items[offset:end]
	}

	var configMapBriefList []modelK8sConfig.ConfigMapBrief
	// 处理configMap数据
	for _, cm := range configMapList.Items {
		var configMapBrief modelK8sConfig.ConfigMapBrief
		configMapBrief.Name = cm.Name
		configMapBrief.Namespace = cm.Namespace
		configMapBrief.CreationTimestamp = cm.CreationTimestamp.Time
		// append
		configMapBriefList = append(configMapBriefList, configMapBrief)
	}

	return configMapBriefList, total, nil
}

// GetConfigMapDetail 获取单个configMap详情
func (ca *ConfigMapService) GetConfigMapDetail(namespace string, name string) (*coreV1.ConfigMap, error) {
	detail, err := global.KF_K8S_Client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

// DeleteConfigMap 删除configMap
func (ca *ConfigMapService) DeleteConfigMap(namespace string, name string) error {
	return global.KF_K8S_Client.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
