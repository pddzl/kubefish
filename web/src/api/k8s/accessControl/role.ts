import { request } from "@/utils/service"
import { type commonListReq, type commonReq } from "@/api/k8s/entry"

// 获取role list
export interface RoleBrief {
  name: string
  namespace: string
  creationTimestamp: string
}

interface RoleBriefList extends PageInfo {
  list: RoleBrief[]
  total: number
}

export const getRoleListApi = (data: commonListReq) => {
  return request<ApiResponseData<RoleBriefList>>({
    url: "/k8s/role/getRoleList",
    method: "post",
    data
  })
}

// 获取role详情
interface RoleDetail {
  metadata: {}
  rules: []
}

export const getRoleDetailApi = (data: commonReq) => {
  return request<ApiResponseData<RoleDetail>>({
    url: "/k8s/role/getRoleDetail",
    method: "post",
    data
  })
}

// 删除role
export const deleteRoleApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/role/deleteRole",
    method: "post",
    data
  })
}
