import { request } from "@/utils/service"

export interface PodData {
  name: string
  namespace: string
  status: string
  node: string
  creationTimestamp: string
}

interface PodListData extends PageInfo {
  list: PodData[]
  total: number
}

interface reqData extends PageInfo {
  namespace: string
}

// 获取所有pod（分页）
export function getPodsApi(data: reqData) {
  return request<IApiResponseData<PodListData>>({
    url: "/k8s/pod/getPods",
    method: "post",
    data
  })
}

interface reqPod {
  namespace: string
  pod: string
}

// 获取pod详情
export function getPodDetailApi(data: reqPod) {
  return request<IApiResponseData<any>>({
    url: "/k8s/pod/getPodDetail",
    method: "post",
    data
  })
}

interface reqPodLog extends reqPod {
  container?: string
  lines: number
  follow?: boolean
}

// 获取pod日志
export function getPodLogApi(data: reqPodLog) {
  return request<IApiResponseData<string>>({
    url: "/k8s/pod/getPodLog",
    method: "post",
    data
  })
}

// 删除
export function deletePodApi(data: reqPod) {
  return request<IApiResponseData<null>>({
    url: "/k8s/pod/deletePod",
    method: "post",
    data
  })
}
