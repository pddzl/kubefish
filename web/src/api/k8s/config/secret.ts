import { request } from "@/utils/service"
import { type commonListReq, type commonReq } from "@/api/k8s/entry"

// 获取secret list
export interface SecretBrief {
  name: string
  namespace: string
  type: string
  creationTimestamp: string
}

interface SecretBriefList extends PageInfo {
  list: SecretBrief[]
  total: number
}

export const getSecretListApi = (data: commonListReq) => {
  return request<ApiResponseData<SecretBriefList>>({
    url: "/k8s/secret/getSecretList",
    method: "post",
    data
  })
}

// 获取secret详情
interface SecretDetail {
  metadata: {}
  data: {}
}

export const getSecretDetailApi = (data: commonReq) => {
  return request<ApiResponseData<SecretDetail>>({
    url: "/k8s/secret/getSecretDetail",
    method: "post",
    data
  })
}

// 删除secret
export const deleteSecretApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/secret/deleteSecret",
    method: "post",
    data
  })
}
