package accessControl

import (
	"context"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8sAccessControl "github.com/pddzl/kubefish/server/model/k8s/accessControl"
)

type RoleService struct{}

// GetRoleList 获取所有role
func (ra *RoleService) GetRoleList(namespace string, pageInfo request.PageInfo) ([]modelK8sAccessControl.RoleBrief, int, error) {
	// 获取roles list
	list, err := global.KF_K8S_Client.RbacV1().Roles(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var roleList rbacV1.RoleList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		roleList.Items = list.Items[offset:]
	} else {
		roleList.Items = list.Items[offset:end]
	}

	// 处理数据
	var roleBriefList []modelK8sAccessControl.RoleBrief
	for _, role := range roleList.Items {
		var roleBrief modelK8sAccessControl.RoleBrief
		roleBrief.Name = role.Name
		roleBrief.Namespace = role.Namespace
		roleBrief.CreationTimestamp = role.CreationTimestamp.Time
		// append
		roleBriefList = append(roleBriefList, roleBrief)
	}

	return roleBriefList, total, nil
}

// GetRoleDetail 获取role详情
func (ra *RoleService) GetRoleDetail(namespace string, name string) (*rbacV1.Role, error) {
	detail, err := global.KF_K8S_Client.RbacV1().Roles(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return detail, nil
}

// DeleteRole 删除role
func (ra *RoleService) DeleteRole(namespace string, name string) error {
	return global.KF_K8S_Client.RbacV1().Roles(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
