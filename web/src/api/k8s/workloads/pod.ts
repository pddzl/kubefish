import { request } from "@/utils/service"

export interface PodBrief {
  name: string
  namespace: string
  status: string
  node: string
  podIP: string
  creationTimestamp: string
}

export interface PodBriefList {
  list: PodBrief[]
  total: number
}

interface reqData extends PageInfo {
  namespace: string
}

// 获取所有pod（分页）
export function getPodsApi(data: reqData) {
  return request<ApiResponseData<PodBriefList>>({
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
  return request<ApiResponseData<any>>({
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
  return request<ApiResponseData<string>>({
    url: "/k8s/pod/getPodLog",
    method: "post",
    data
  })
}

// 删除
export function deletePodApi(data: reqPod) {
  return request<ApiResponseData<null>>({
    url: "/k8s/pod/deletePod",
    method: "post",
    data
  })
}
