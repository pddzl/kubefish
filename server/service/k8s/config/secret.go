package config

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8sConfig "github.com/pddzl/kubefish/server/model/k8s/config"
)

type SecretService struct{}

// GetSecretList 获取所有secret
func (ss *SecretService) GetSecretList(namespace string, info request.PageInfo) ([]modelK8sConfig.SecretBrief, int, error) {
	// 获取secret list
	list, err := global.KF_K8S_Client.CoreV1().Secrets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var secretList coreV1.SecretList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		secretList.Items = list.Items[offset:]
	} else {
		secretList.Items = list.Items[offset:end]
	}

	var secretBriefList []modelK8sConfig.SecretBrief
	for _, secret := range secretList.Items {
		var secretBrief modelK8sConfig.SecretBrief
		secretBrief.Name = secret.Name
		secretBrief.Namespace = secret.Namespace
		secretBrief.Type = string(secret.Type)
		secretBrief.CreationTimestamp = secret.CreationTimestamp.Time
		// append
		secretBriefList = append(secretBriefList, secretBrief)
	}

	return secretBriefList, total, nil
}

// GetSecretDetail 获取单个secret详情
func (ss *SecretService) GetSecretDetail(namespace string, name string) (*coreV1.Secret, error) {
	detail, err := global.KF_K8S_Client.CoreV1().Secrets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

// DeleteSecret 删除secret
func (ss *SecretService) DeleteSecret(namespace string, name string) error {
	return global.KF_K8S_Client.CoreV1().Secrets(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
