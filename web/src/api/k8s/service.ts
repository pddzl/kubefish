import { request } from "@/utils/service"
import { type commonListReq, type commonReq } from "./entry"

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

// 删除service
export const deleteServiceApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/service/deleteService",
    method: "post",
    data
  })
}
