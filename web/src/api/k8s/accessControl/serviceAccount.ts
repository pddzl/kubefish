import { request } from "@/utils/service"
import { type commonListReq, type commonReq } from "@/api/k8s/entry"

// 获取serviceAccount list
export interface ServiceAccountBrief {
  name: string
  namespace: string
  creationTimestamp: string
}

interface ServiceAccountBriefList extends PageInfo {
  list: ServiceAccountBrief[]
  total: number
}

export const getServiceAccountListApi = (data: commonListReq) => {
  return request<ApiResponseData<ServiceAccountBriefList>>({
    url: "/k8s/serviceAccount/getServiceAccountList",
    method: "post",
    data
  })
}

// 获取serviceAccount详情
export interface Secret {
  name: string
}

interface ServiceAccountDetail {
  metadata: {}
  secrets: Secret[]
}

export const getServiceAccountDetailApi = (data: commonReq) => {
  return request<ApiResponseData<ServiceAccountDetail>>({
    url: "/k8s/serviceAccount/getServiceAccountDetail",
    method: "post",
    data
  })
}

// 删除serviceAccount
export const deleteServiceAccountApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/serviceAccount/deleteServiceAccount",
    method: "post",
    data
  })
}
