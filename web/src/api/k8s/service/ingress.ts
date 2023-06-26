import { request } from "@/utils/service"
import { type commonListReq, type commonReq } from "../entry"

export interface IngressBrief {
  name: string
  namespace: string
  ip: string
  hosts: string[]
  creationTimestamp: string
}

interface IngressBriefList extends PageInfo {
  list: IngressBrief[]
  total: number
}

// 获取ingress list
export const getIngressListApi = (data: commonListReq) => {
  return request<ApiResponseData<IngressBriefList>>({
    url: "/k8s/ingress/getIngressList",
    method: "post",
    data
  })
}

// 获取ingress 详情
export const getIngressDetailApi = (data: commonReq) => {
  return request<ApiResponseData<any>>({
    url: "/k8s/ingress/getIngressDetail",
    method: "post",
    data
  })
}

// 删除ingress
export const deleteIngressApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/ingress/deleteIngress",
    method: "post",
    data
  })
}
