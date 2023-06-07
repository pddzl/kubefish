package k8s

import (
	"context"
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8s "github.com/pddzl/kubefish/server/model/k8s"
)

type DaemonSetService struct{}

// GetDaemonSets 获取所有daemonSet
func (ds *DaemonSetService) GetDaemonSets(namespace string, label string, field string, pageInfo request.PageInfo) ([]modelK8s.DaemonSetBrief, int, error) {
	// 获取daemonSet list数据
	var option metaV1.ListOptions
	if label != "" {
		option.LabelSelector = label
	}
	if field != "" {
		option.FieldSelector = field
	}

	list, err := global.KF_K8S_Client.AppsV1().DaemonSets(namespace).List(context.TODO(), option)
	if err != nil {
		return nil, 0, err
	}

	var daemonSetList appsV1.DaemonSetList
	// 分页
	end := pageInfo.PageSize * pageInfo.Page
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		daemonSetList.Items = list.Items[offset:]
	} else {
		daemonSetList.Items = list.Items[offset:end]
	}

	var daemonSetBriefList []modelK8s.DaemonSetBrief
	// 处理daemonSets数据
	for _, dsV := range daemonSetList.Items {
		var daemonSetBrief modelK8s.DaemonSetBrief
		daemonSetBrief.Name = dsV.Name
		daemonSetBrief.NameSpace = dsV.Namespace
		daemonSetBrief.Pods = fmt.Sprintf("%d / %d", dsV.Status.NumberAvailable, dsV.Status.DesiredNumberScheduled)
		daemonSetBrief.CreationTimestamp = dsV.CreationTimestamp.Time
		// append
		daemonSetBriefList = append(daemonSetBriefList, daemonSetBrief)
	}

	return daemonSetBriefList, total, nil
}

// GetDaemonSetDetail 获取daemonSet 详情
func (ds *DaemonSetService) GetDaemonSetDetail(namespace string, name string) (*modelK8s.DaemonSetDetail, error) {
	// 获取daemonSet原生数据
	daemonSet, err := global.KF_K8S_Client.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理daemonSet数据
	var daemonSetDetail modelK8s.DaemonSetDetail
	// metadata
	daemonSetDetail.ObjectMeta = modelK8s.NewObjectMeta(daemonSet.ObjectMeta)
	// spec
	daemonSetDetail.Spec.Selector = daemonSet.Spec.Selector.MatchLabels
	daemonSetDetail.Spec.UpdateStrategy.Type = string(daemonSet.Spec.UpdateStrategy.Type)
	daemonSetDetail.Spec.UpdateStrategy.RollingUpdate.MaxSurge = daemonSet.Spec.UpdateStrategy.RollingUpdate.MaxSurge
	daemonSetDetail.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable = daemonSet.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable
	daemonSetDetail.Spec.MinReadySeconds = daemonSet.Spec.MinReadySeconds
	daemonSetDetail.Spec.RevisionHistoryLimit = daemonSet.Spec.RevisionHistoryLimit
	// status
	daemonSetDetail.Status.CurrentNumberScheduled = daemonSet.Status.CurrentNumberScheduled
	daemonSetDetail.Status.NumberMisscheduled = daemonSet.Status.NumberMisscheduled
	daemonSetDetail.Status.DesiredNumberScheduled = daemonSet.Status.DesiredNumberScheduled
	daemonSetDetail.Status.NumberReady = daemonSet.Status.NumberReady
	daemonSetDetail.Status.UpdatedNumberScheduled = daemonSet.Status.UpdatedNumberScheduled
	daemonSetDetail.Status.NumberAvailable = daemonSet.Status.NumberAvailable
	daemonSetDetail.Status.NumberUnavailable = daemonSet.Status.NumberUnavailable
	// status -> conditions
	for _, condition := range daemonSet.Status.Conditions {
		var daemonSetCondition modelK8s.DaemonSetCondition
		daemonSetCondition.Type = string(condition.Type)
		daemonSetCondition.Status = string(condition.Status)
		daemonSetCondition.LastTransitionTime = condition.LastTransitionTime.Time
		daemonSetCondition.Reason = condition.Reason
		daemonSetCondition.Message = condition.Message
		// append
		daemonSetDetail.Status.Conditions = append(daemonSetDetail.Status.Conditions, daemonSetCondition)
	}

	return &daemonSetDetail, nil
}

// GetDaemonSetPods 获取daemonSet关联的pod
func (ds *DaemonSetService) GetDaemonSetPods(namespace string, name string, pageInfo request.PageInfo) ([]modelK8s.PodBrief, int, error) {
	// 获取daemonSet原始数据
	rSet, err := global.KF_K8S_Client.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, 0, err
	}

	labelSelector := labels.Set(rSet.Spec.Selector.MatchLabels).String()
	var hostService PodService
	podBriefList, total, err := hostService.GetPods(namespace, labelSelector, "", pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		global.KF_LOG.Error("获取daemonSet关联的pod失败")
		return nil, 0, err
	} else {
		return podBriefList, total, nil
	}
}

// DeleteDaemonSet 删除daemonSet
func (ds *DaemonSetService) DeleteDaemonSet(namespace string, name string) error {
	return global.KF_K8S_Client.AppsV1().DaemonSets(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
