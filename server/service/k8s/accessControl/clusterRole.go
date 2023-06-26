package accessControl

import (
	"context"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8sAccessControl "github.com/pddzl/kubefish/server/model/k8s/accessControl"
)

type ClusterRoleService struct{}

// GetClusterRoleList 获取所有clusterRole
func (crs *ClusterRoleService) GetClusterRoleList(pageInfo request.PageInfo) ([]modelK8sAccessControl.ClusterRoleBrief, int, error) {
	// 获取clusterRole list
	list, err := global.KF_K8S_Client.RbacV1().ClusterRoles().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var clusterRoleList rbacV1.ClusterRoleList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		clusterRoleList.Items = list.Items[offset:]
	} else {
		clusterRoleList.Items = list.Items[offset:end]
	}

	// 处理数据
	var clusterRoleBriefList []modelK8sAccessControl.ClusterRoleBrief
	for _, clusterRole := range clusterRoleList.Items {
		var clusterRoleBrief modelK8sAccessControl.ClusterRoleBrief
		clusterRoleBrief.Name = clusterRole.Name
		clusterRoleBrief.CreationTimestamp = clusterRole.CreationTimestamp.Time
		// append
		clusterRoleBriefList = append(clusterRoleBriefList, clusterRoleBrief)
	}

	return clusterRoleBriefList, total, nil
}

// GetClusterRoleDetail 获取clusterRole详情
func (crs *ClusterRoleService) GetClusterRoleDetail(name string) (*rbacV1.ClusterRole, error) {
	detail, err := global.KF_K8S_Client.RbacV1().ClusterRoles().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

// DeleteClusterRole 删除clusterRole
func (crs *ClusterRoleService) DeleteClusterRole(name string) error {
	return global.KF_K8S_Client.RbacV1().ClusterRoles().Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
