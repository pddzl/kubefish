import { request } from "@/utils/service"
import { type commonListReq, type commonReq } from "@/api/k8s/entry"

// 获取roleBinding list
export interface RoleBindingBrief {
  name: string
  namespace: string
  creationTimestamp: string
}

interface RoleBindingBriefList extends PageInfo {
  list: RoleBindingBrief[]
  total: number
}

export const getRoleBindingListApi = (data: commonListReq) => {
  return request<ApiResponseData<RoleBindingBriefList>>({
    url: "/k8s/roleBinding/getRoleBindingList",
    method: "post",
    data
  })
}

// 获取roleBinding详情
interface RoleBindingDetail {
  metadata: {}
  subjects: []
  roleRef: { name: string; kind: string; apiGroup: string }
}

export const getRoleBindingDetailApi = (data: commonReq) => {
  return request<ApiResponseData<RoleBindingDetail>>({
    url: "/k8s/roleBinding/getRoleBindingDetail",
    method: "post",
    data
  })
}

// 删除roleBinding
export const deleteRoleBindingApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/roleBinding/deleteRoleBinding",
    method: "post",
    data
  })
}
