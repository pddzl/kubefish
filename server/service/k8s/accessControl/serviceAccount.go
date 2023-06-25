package accessControl

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8sAccessControl "github.com/pddzl/kubefish/server/model/k8s/accessControl"
)

type ServiceAccountService struct{}

// GetServiceAccountList 获取所有serviceAccount
func (sas *ServiceAccountService) GetServiceAccountList(namespace string, pageInfo request.PageInfo) ([]modelK8sAccessControl.ServiceAccountBrief, int, error) {
	// 获取serviceAccount List
	list, err := global.KF_K8S_Client.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var serviceAccountList coreV1.ServiceAccountList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		serviceAccountList.Items = list.Items[offset:]
	} else {
		serviceAccountList.Items = list.Items[offset:end]
	}

	var serviceAccountBriefList []modelK8sAccessControl.ServiceAccountBrief
	// 处理list数据
	for _, serviceAccount := range serviceAccountList.Items {
		var serviceAccountBrief modelK8sAccessControl.ServiceAccountBrief
		serviceAccountBrief.Name = serviceAccount.Name
		serviceAccountBrief.Namespace = serviceAccount.Namespace
		serviceAccountBrief.CreationTimestamp = serviceAccount.CreationTimestamp.Time
		// append
		serviceAccountBriefList = append(serviceAccountBriefList, serviceAccountBrief)
	}

	return serviceAccountBriefList, total, nil
}

// GetServiceAccountDetail 获取serviceAccount详情
func (sas *ServiceAccountService) GetServiceAccountDetail(namespace string, name string) (*coreV1.ServiceAccount, error) {
	// 获取原生数据
	detail, err := global.KF_K8S_Client.CoreV1().ServiceAccounts(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

// DeleteServiceAccount 删除serviceAccount
func (sas *ServiceAccountService) DeleteServiceAccount(namespace string, name string) error {
	return global.KF_K8S_Client.CoreV1().ServiceAccounts(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
