import { request } from "@/utils/service"
import { PodBriefList } from "../workloads/pod"

export interface NodeData {
  name: string
  internalIP: string
  role: string[]
  status: string
  cpu: string
  memory: string
  creationTimestamp: string
}

export interface NodeDataPageInfo {
  list: NodeData[]
  total: number
  page: number
  pageSize: number
}

// 获取所有节点（分页）
export function getNodesApi(data: PageInfo) {
  return request<ApiResponseData<NodeDataPageInfo>>({
    url: "/k8s/node/getNodes",
    method: "post",
    data
  })
}

// export interface NodeDetail {
//   metadata: {
//     name: string
//     label: object
//     annotations: object
//     creationTimestamp: string
//     uid: string
//   }
//   spec: {
//     podCIDR: string
//     taints: object[]
//   }
//   status: {
//     conditions: object[]
//     address: object[]
//     port: {
//       kubeletEndpoint: {
//         Port: number
//       }
//     }
//     nodeInfo: object
//   }
//   allocatedResources: object
// }

// 获取某个节点详情
export function getNodeDetailApi(params: { name: string }) {
  return request<ApiResponseData<any>>({
    url: "/k8s/node/getNodeDetail",
    method: "get",
    params
  })
}

interface nodePodReq extends PageInfo {
  node_name: string
}

export function getNodePodsApi(data: nodePodReq) {
  return request<ApiResponseData<PodBriefList>>({
    url: "/k8s/node/getNodePods",
    method: "post",
    data
  })
}
