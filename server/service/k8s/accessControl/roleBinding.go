package accessControl

import (
	"context"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8sAccessControl "github.com/pddzl/kubefish/server/model/k8s/accessControl"
)

type RoleBindingService struct{}

// GetRoleBindingList 获取所有roleBinding
func (rba *RoleBindingService) GetRoleBindingList(namespace string, pageInfo request.PageInfo) ([]modelK8sAccessControl.RoleBindingBrief, int, error) {
	// 获取roleBinding list
	list, err := global.KF_K8S_Client.RbacV1().RoleBindings(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var roleBindingList rbacV1.RoleBindingList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		roleBindingList.Items = list.Items[offset:]
	} else {
		roleBindingList.Items = list.Items[offset:end]
	}

	var roleBindingBriefList []modelK8sAccessControl.RoleBindingBrief
	// 处理数据
	for _, roleBinding := range roleBindingList.Items {
		var roleBindingBrief modelK8sAccessControl.RoleBindingBrief
		roleBindingBrief.Name = roleBinding.Name
		roleBindingBrief.Namespace = roleBinding.Namespace
		roleBindingBrief.CreationTimestamp = roleBinding.CreationTimestamp.Time
		// append
		roleBindingBriefList = append(roleBindingBriefList, roleBindingBrief)
	}

	return roleBindingBriefList, total, nil
}

// GetRoleBindingDetail 获取roleBinding详情
func (rba *RoleBindingService) GetRoleBindingDetail(namespace string, name string) (*rbacV1.RoleBinding, error) {
	detail, err := global.KF_K8S_Client.RbacV1().RoleBindings(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

// DeleteRoleBinding 删除roleBinding
func (rba *RoleBindingService) DeleteRoleBinding(namespace string, name string) error {
	return global.KF_K8S_Client.RbacV1().RoleBindings(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
