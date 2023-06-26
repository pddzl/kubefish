package accessControl

import (
	"context"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8sAccessControl "github.com/pddzl/kubefish/server/model/k8s/accessControl"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterRoleBindingService struct{}

// GetClusterRoleBindingList 获取所有clusterRoleBinding
func (c *ClusterRoleBindingService) GetClusterRoleBindingList(pageInfo request.PageInfo) ([]modelK8sAccessControl.ClusterRoleBindingBrief, int, error) {
	list, err := global.KF_K8S_Client.RbacV1().ClusterRoleBindings().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var clusterRoleBindingList rbacV1.ClusterRoleBindingList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		clusterRoleBindingList.Items = list.Items[offset:]
	} else {
		clusterRoleBindingList.Items = list.Items[offset:end]
	}

	var clusterRoleBindingBriefList []modelK8sAccessControl.ClusterRoleBindingBrief
	// 处理数据
	for _, role := range clusterRoleBindingList.Items {
		var clusterRoleBindingBrief modelK8sAccessControl.ClusterRoleBindingBrief
		clusterRoleBindingBrief.Name = role.Name
		clusterRoleBindingBrief.CreationTimestamp = role.CreationTimestamp.Time
		// append
		clusterRoleBindingBriefList = append(clusterRoleBindingBriefList, clusterRoleBindingBrief)
	}

	return clusterRoleBindingBriefList, total, nil
}

// GetClusterRoleBinding 获取clusterRoleBinding详情
func (c *ClusterRoleBindingService) GetClusterRoleBinding(name string) (*rbacV1.ClusterRoleBinding, error) {
	detail, err := global.KF_K8S_Client.RbacV1().ClusterRoleBindings().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

// DeleteClusterRoleBinding 删除clusterRoleBinding
func (c *ClusterRoleBindingService) DeleteClusterRoleBinding(name string) error {
	return global.KF_K8S_Client.RbacV1().ClusterRoleBindings().Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
