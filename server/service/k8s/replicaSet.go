package k8s

import (
	"context"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8s "github.com/pddzl/kubefish/server/model/k8s"
)

type ReplicaSetService struct{}

// GetReplicaSets 获取replicaSet列表
func (rs *ReplicaSetService) GetReplicaSets(namespace string, info request.PageInfo) ([]modelK8s.ReplicaSetBrief, int, error) {
	// 获取replicaSet list
	list, err := global.KF_K8S_Client.AppsV1().ReplicaSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var replicaSetList appsV1.ReplicaSetList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		replicaSetList.Items = list.Items[offset:]
	} else {
		replicaSetList.Items = list.Items[offset:end]
	}

	// 处理replicaSet数据
	var replicaSetBriefList []modelK8s.ReplicaSetBrief
	for _, rsV := range replicaSetList.Items {
		var replicaSetBrief modelK8s.ReplicaSetBrief
		replicaSetBrief.Name = rsV.Name
		replicaSetBrief.NameSpace = rsV.Namespace
		replicaSetBrief.AvailableReplicas = rsV.Status.AvailableReplicas
		replicaSetBrief.Replicas = rsV.Status.Replicas
		replicaSetBrief.CreationTimestamp = rsV.CreationTimestamp.Time
		// append
		replicaSetBriefList = append(replicaSetBriefList, replicaSetBrief)
	}

	return replicaSetBriefList, total, nil
}

// GetReplicaSetDetail 获取replicaSet详情
func (rs *ReplicaSetService) GetReplicaSetDetail(namespace string, replicaSetName string) (*modelK8s.ReplicaSetDetail, error) {
	replicaSet, err := global.KF_K8S_Client.AppsV1().ReplicaSets(namespace).Get(context.TODO(), replicaSetName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理replicaSet数据
	var replicaSetDetail modelK8s.ReplicaSetDetail
	// metadata
	replicaSetDetail.ObjectMeta = modelK8s.NewObjectMeta(replicaSet.ObjectMeta)
	// spec
	replicaSetDetail.Spec.Replicas = replicaSet.Spec.Replicas
	replicaSetDetail.Spec.MinReadySeconds = replicaSet.Spec.MinReadySeconds
	replicaSetDetail.Spec.Selector = replicaSet.Spec.Selector.MatchLabels
	// status
	replicaSetDetail.Status.Replicas = replicaSet.Status.Replicas
	replicaSetDetail.Status.FullyLabeledReplicas = replicaSet.Status.FullyLabeledReplicas
	replicaSetDetail.Status.ReadyReplicas = replicaSet.Status.ReadyReplicas
	replicaSetDetail.Status.AvailableReplicas = replicaSet.Status.AvailableReplicas
	// status -> condition
	for _, condition := range replicaSet.Status.Conditions {
		var replicaSetCondition modelK8s.ReplicaSetCondition
		replicaSetCondition.Type = string(condition.Type)
		replicaSetCondition.Status = string(condition.Status)
		replicaSetCondition.LastTransitionTime = condition.LastTransitionTime.Time
		replicaSetCondition.Reason = condition.Reason
		replicaSetCondition.Message = condition.Message
		// append
		replicaSetDetail.Status.Conditions = append(replicaSetDetail.Status.Conditions, replicaSetCondition)
	}

	return &replicaSetDetail, nil
}
