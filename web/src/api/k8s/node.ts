import { request } from "@/utils/service"

// type status = "True" | "Unknown"

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
export function getNodeDetail(params: { name: string }) {
  return request<ApiResponseData<any>>({
    url: "/k8s/node/getNodeDetail",
    method: "get",
    params
  })
}

export interface NodePods {
  metadata: {
    name: string
    namespace: string
    creationTimestamp: string
  }
  status: string
  nodeName: string
}

export interface NodePodsPageInfo {
  list: NodePods[]
  total: number
  page: number
  pageSize: number
}

interface nodePodsData extends PageInfo {
  node_name: string
}

export function getNodePods(data: nodePodsData) {
  return request<ApiResponseData<NodePodsPageInfo>>({
    url: "/k8s/node/getNodePods",
    method: "post",
    data
  })
}
