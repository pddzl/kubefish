import { request } from "@/utils/service"
import { type nameReq } from "@/api/k8s/entry"

// 获取ClusterRoleBinding list
export interface ClusterRoleBindingBrief {
  name: string
  creationTimestamp: string
}

interface ClusterRoleBindingBriefList extends PageInfo {
  list: ClusterRoleBindingBrief[]
  total: number
}

export const getClusterRoleBindingListApi = (data: PageInfo) => {
  return request<ApiResponseData<ClusterRoleBindingBriefList>>({
    url: "/k8s/clusterRoleBinding/getClusterRoleBindingList",
    method: "post",
    data
  })
}

// 获取ClusterRoleBinding详情
interface ClusterRoleBindingDetail {
  metadata: {}
  subjects: []
  roleRef: { name: string; kind: string; apiGroup: string }
}

export const getClusterRoleBindingDetailApi = (data: nameReq) => {
  return request<ApiResponseData<ClusterRoleBindingDetail>>({
    url: "/k8s/clusterRoleBinding/getClusterRoleBindingDetail",
    method: "post",
    data
  })
}

// 删除ClusterRoleBinding
export const deleteClusterRoleBindingApi = (data: nameReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/clusterRoleBinding/deleteClusterRoleBinding",
    method: "post",
    data
  })
}
