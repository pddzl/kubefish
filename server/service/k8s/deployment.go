package k8s

import (
	"context"
	"fmt"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8s "github.com/pddzl/kubefish/server/model/k8s"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type DeploymentService struct{}

// GetDeploymentList 获取所有deployment
func (ds *DeploymentService) GetDeploymentList(namespace string, label string, field string, info request.PageInfo) ([]modelK8s.DeploymentBrief, int, error) {
	// 获取deployment list
	var option metaV1.ListOptions
	if label != "" {
		option.LabelSelector = label
	}
	if field != "" {
		option.FieldSelector = field
	}

	list, err := global.KF_K8S_Client.AppsV1().Deployments(namespace).List(context.TODO(), option)
	if err != nil {
		return nil, 0, err
	}

	var deploymentList appsV1.DeploymentList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		deploymentList.Items = list.Items[offset:]
	} else {
		deploymentList.Items = list.Items[offset:end]
	}

	var deploymentBriefList []modelK8s.DeploymentBrief
	// 处理list数据
	for _, dm := range deploymentList.Items {
		var deploymentBrief modelK8s.DeploymentBrief
		deploymentBrief.Name = dm.Name
		deploymentBrief.Namespace = dm.Namespace
		deploymentBrief.AvailableReplicas = dm.Status.AvailableReplicas
		deploymentBrief.Replicas = dm.Status.Replicas
		deploymentBrief.CreationTimestamp = dm.CreationTimestamp.Time
		// append
		deploymentBriefList = append(deploymentBriefList, deploymentBrief)
	}

	return deploymentBriefList, total, nil
}

// GetDeploymentDetail 获取deployment详情
func (ds *DeploymentService) GetDeploymentDetail(namespace string, name string) (*modelK8s.DeploymentDetail, error) {
	deployment, err := global.KF_K8S_Client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理deployment数据
	var deploymentDetail modelK8s.DeploymentDetail
	// metadata
	deploymentDetail.ObjectMeta = modelK8s.NewObjectMeta(deployment.ObjectMeta)
	// spec
	deploymentDetail.Spec.Replicas = deployment.Spec.Replicas
	deploymentDetail.Spec.Selector = deployment.Spec.Selector.MatchLabels
	deploymentDetail.Spec.Strategy.DeploymentStrategyType = string(deployment.Spec.Strategy.Type)
	deploymentDetail.Spec.Strategy.RollingUpdateDeployment.MaxUnavailable = deployment.Spec.Strategy.RollingUpdate.MaxUnavailable
	deploymentDetail.Spec.Strategy.RollingUpdateDeployment.MaxSurge = deployment.Spec.Strategy.RollingUpdate.MaxSurge
	deploymentDetail.Spec.MinReadySeconds = deployment.Spec.MinReadySeconds
	deploymentDetail.Spec.RevisionHistoryLimit = deployment.Spec.RevisionHistoryLimit
	// status
	deploymentDetail.Status.Replicas = deployment.Status.Replicas
	deploymentDetail.Status.UpdatedReplicas = deployment.Status.UpdatedReplicas
	deploymentDetail.Status.ReadyReplicas = deployment.Status.ReadyReplicas
	deploymentDetail.Status.AvailableReplicas = deployment.Status.AvailableReplicas
	deploymentDetail.Status.UnavailableReplicas = deployment.Status.UnavailableReplicas
	// status -> condition
	for _, dsc := range deployment.Status.Conditions {
		var condition modelK8s.DeploymentCondition
		condition.Type = string(dsc.Type)
		condition.Status = string(dsc.Status)
		condition.LastTransitionTime = dsc.LastTransitionTime.Time
		condition.LastUpdateTime = dsc.LastUpdateTime.Time
		condition.Reason = dsc.Reason
		condition.Message = dsc.Message
		// append
		deploymentDetail.Status.Conditions = append(deploymentDetail.Status.Conditions, condition)
	}

	return &deploymentDetail, nil
}

// DeleteDeployment 删除deployment
func (ds *DeploymentService) DeleteDeployment(namespace string, name string) error {
	return global.KF_K8S_Client.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

// GetDeploymentRs 获取deployment关联的rs
func (ds *DeploymentService) GetDeploymentRs(namespace string, name string, pageInfo request.PageInfo) (*modelK8s.NewReplicaSet, error) {
	deployment, err := global.KF_K8S_Client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	labelSelector := labels.Set(deployment.Spec.Selector.MatchLabels).String()

	// 获取replicaSet
	replicaSetList, err := global.KF_K8S_Client.AppsV1().ReplicaSets(namespace).List(context.TODO(), metaV1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return nil, err
	}

	if len(replicaSetList.Items) == 0 {
		return nil, nil
	}

	var newReplicaSet modelK8s.NewReplicaSet
	newReplicaSet.Name = replicaSetList.Items[0].Name
	newReplicaSet.NameSpace = replicaSetList.Items[0].Namespace
	newReplicaSet.Labels = replicaSetList.Items[0].Labels
	newReplicaSet.CreationTimestamp = replicaSetList.Items[0].CreationTimestamp.Time
	newReplicaSet.Replicas = fmt.Sprintf("%d / %d", replicaSetList.Items[0].Status.AvailableReplicas, replicaSetList.Items[0].Status.Replicas)

	return &newReplicaSet, nil

}
