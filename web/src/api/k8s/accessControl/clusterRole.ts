import { request } from "@/utils/service"
import { type nameReq } from "@/api/k8s/entry"

// 获取clusterRole list
export interface ClusterRoleBrief {
  name: string
  creationTimestamp: string
}

interface ClusterRoleBriefList extends PageInfo {
  list: ClusterRoleBrief[]
  total: number
}

export const getClusterRoleListApi = (data: PageInfo) => {
  return request<ApiResponseData<ClusterRoleBriefList>>({
    url: "/k8s/clusterRole/getClusterRoleList",
    method: "post",
    data
  })
}

// 获取clusterRole详情
interface ClusterRoleDetail {
  metadata: {}
  rules: []
}

export const getClusterRoleDetailApi = (data: nameReq) => {
  return request<ApiResponseData<ClusterRoleDetail>>({
    url: "/k8s/clusterRole/getClusterRoleDetail",
    method: "post",
    data
  })
}

// 删除clusterRole
export const deleteClusterRoleApi = (data: nameReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/clusterRole/deleteClusterRole",
    method: "post",
    data
  })
}
