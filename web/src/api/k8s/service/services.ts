import { request } from "@/utils/service"
import { type commonListReq, type commonRelatedReq, type commonReq } from "../entry"
import { type PodBriefList } from "../workloads/pod"

export interface ServiceBrief {
  name: string
  namespace: string
  clusterIP: string
  type: string
  external: string
  creationTimestamp: string
}

interface ServiceBriefList extends PageInfo {
  list: ServiceBrief[]
  total: number
}

// 获取service list
export const getServicesApi = (data: commonListReq) => {
  return request<ApiResponseData<ServiceBriefList>>({
    url: "/k8s/service/getServices",
    method: "post",
    data
  })
}

// 获取service 详情
export const getServiceDetailApi = (data: commonReq) => {
  return request<ApiResponseData<any>>({
    url: "/k8s/service/getServiceDetail",
    method: "post",
    data
  })
}

// 获取service关联的pod
export const getServicePodsApi = (data: commonRelatedReq) => {
  return request<ApiResponseData<PodBriefList>>({
    url: "/k8s/service/getServicePods",
    method: "post",
    data
  })
}

// 删除service
export const deleteServiceApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/service/deleteService",
    method: "post",
    data
  })
}
