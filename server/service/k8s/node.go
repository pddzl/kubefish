package k8s

import (
	"context"
	"fmt"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/request"
	modelK8s "github.com/pddzl/kubefish/server/model/k8s"
	k8sResp "github.com/pddzl/kubefish/server/model/k8s/response"
	"go.uber.org/zap"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"regexp"
	"strconv"
)

type NodeService struct{}

func (n *NodeService) GetNodes(info request.PageInfo) ([]k8sResp.NodeBrief, int, error) {
	// 获取node list
	list, err := global.KF_K8S_Client.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var nodeList coreV1.NodeList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		nodeList.Items = list.Items[offset:]
	} else {
		nodeList.Items = list.Items[offset:end]
	}

	var nodeBriefList []k8sResp.NodeBrief
	// 处理list数据
	roleRe, _ := regexp.Compile("node-role.kubernetes.io/(.*)")
	for _, node := range nodeList.Items {
		var nodeBrief k8sResp.NodeBrief
		nodeBrief.Name = node.Name
		nodeBrief.InternalIP = node.Status.Addresses[0].Address
		for label := range node.ObjectMeta.Labels {
			role := roleRe.FindStringSubmatch(label)
			if len(role) == 2 {
				nodeBrief.Roles = append(nodeBrief.Roles, role[1])
			}
		}
		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" {
				nodeBrief.Status = string(condition.Status)
			}
		}
		nodeBrief.Cpu = node.Status.Capacity.Cpu().String()
		nodeBrief.Memory = node.Status.Capacity.Memory().String()
		nodeBrief.CreationTimestamp = node.CreationTimestamp.Time
		nodeBriefList = append(nodeBriefList, nodeBrief)
	}

	return nodeBriefList, total, nil
}

func (n *NodeService) GetNodeDetail(name string) (*modelK8s.NodeDetail, error) {
	node, err := global.KF_K8S_Client.CoreV1().Nodes().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		global.KF_LOG.Error("获取节点信息失败", zap.Error(err))
		return nil, err
	}

	// 处理node数据
	var nodeDetail modelK8s.NodeDetail
	// metadata
	nodeDetail.Metadata = modelK8s.NewObjectMeta(node.ObjectMeta)
	// spec
	nodeDetail.Spec.PodCIDR = node.Spec.PodCIDR
	nodeDetail.Spec.Unschedulable = node.Spec.Unschedulable
	nodeDetail.Spec.Taints = node.Spec.Taints
	// status
	nodeDetail.Status.Phase = string(node.Status.Phase)
	nodeDetail.Status.Conditions = node.Status.Conditions
	nodeDetail.Status.Addresses = node.Status.Addresses
	nodeDetail.Status.DaemonEndpoints = node.Status.DaemonEndpoints
	nodeDetail.Status.NodeInfo = node.Status.NodeInfo
	nodeDetail.Status.Images = node.Status.Images

	// 获取节点可分配的资源总量
	nodeDetail.NodeAllocatedResources.AllocatedCPU = node.Status.Allocatable.Cpu().MilliValue()
	nodeDetail.NodeAllocatedResources.AllocatedMemory = node.Status.Allocatable.Memory().MilliValue()
	nodeDetail.NodeAllocatedResources.AllocatedPods = node.Status.Capacity.Pods().Value()

	// 获取节点上所有运行的 Pod 的 CPU Request 和占比
	pods, err := getNodePods(name)
	if err != nil {
		global.KF_LOG.Error("获取节点pod失败", zap.Error(err))
	}

	var cpuRequest, cpuLimit int64
	var memoryRequest, memoryLimit int64
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			// cpu request
			cpuR := container.Resources.Requests.Cpu().MilliValue()
			cpuRequest += cpuR
			// cpu limit
			cpuL := container.Resources.Limits.Cpu().MilliValue()
			cpuLimit += cpuL
			// memory request
			memoryR := container.Resources.Requests.Memory().MilliValue()
			memoryRequest += memoryR
			// memory limit
			memoryL := container.Resources.Limits.Memory().MilliValue()
			memoryLimit += memoryL
		}
	}

	// cpu
	nodeDetail.NodeAllocatedResources.CPURequests = cpuRequest
	if cpuRequest > 0 {
		nodeDetail.NodeAllocatedResources.CPURequestsPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", float64(cpuRequest)/float64(nodeDetail.NodeAllocatedResources.AllocatedCPU)), 64)
	}
	nodeDetail.NodeAllocatedResources.CPULimits = cpuLimit
	if cpuLimit > 0 {
		nodeDetail.NodeAllocatedResources.CPULimitsPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", float64(cpuLimit)/float64(nodeDetail.NodeAllocatedResources.AllocatedCPU)), 64)
	}

	// memory
	nodeDetail.NodeAllocatedResources.MemoryRequests = memoryRequest
	if memoryRequest > 0 {
		nodeDetail.NodeAllocatedResources.MemoryRequestsPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", float64(memoryRequest)/float64(nodeDetail.NodeAllocatedResources.AllocatedMemory)), 64)
	}
	nodeDetail.NodeAllocatedResources.MemoryLimits = memoryLimit
	if memoryLimit > 0 {
		nodeDetail.NodeAllocatedResources.MemoryLimitsPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", float64(memoryLimit)/float64(nodeDetail.NodeAllocatedResources.AllocatedMemory)), 64)
	}

	// pod
	podNum := len(pods.Items)
	nodeDetail.NodeAllocatedResources.PodNum = podNum
	if podNum > 0 {
		nodeDetail.NodeAllocatedResources.PodPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", float64(podNum)/float64(nodeDetail.NodeAllocatedResources.AllocatedPods)), 64)
	}

	return &nodeDetail, nil
}

func getNodePods(name string) (*coreV1.PodList, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + name +
		",status.phase!=" + string(coreV1.PodSucceeded) +
		",status.phase!=" + string(coreV1.PodFailed))

	if err != nil {
		return nil, err
	}

	return global.KF_K8S_Client.CoreV1().Pods(coreV1.NamespaceAll).List(context.TODO(), metaV1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
}

// GetNodePods 获取节点上的pod信息
func (n *NodeService) GetNodePods(name string, info request.PageInfo) ([]modelK8s.PodBrief, int, error) {
	// 制作label
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + name +
		",status.phase!=" + string(coreV1.PodSucceeded) +
		",status.phase!=" + string(coreV1.PodFailed))

	if err != nil {
		return nil, 0, err
	}

	var podService PodService
	podList, total, err := podService.GetPods(coreV1.NamespaceAll, fieldSelector.String(), "field", info.Page, info.PageSize)
	if err != nil {
		global.KF_LOG.Error("获取node关联的pod失败")
		return nil, 0, err
	} else {
		return podList, total, nil
	}
}
