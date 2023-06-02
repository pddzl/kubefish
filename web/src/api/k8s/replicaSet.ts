import { request } from "@/utils/service"
import { reqNamespace } from "./entry"

export interface replicaSetData {
  name: string
  namespace: string
  availableReplicas: number
  replicas: number
  creationTimestamp: string
}

interface replicaSetListData extends PageInfo {
  list: replicaSetData[]
  total: number
}

// 获取replicaSet list
export const getReplicaSetsApi = (data: reqNamespace) => {
  return request<ApiResponseData<replicaSetListData>>({
    url: '/k8s/replicaSet/getReplicaSets',
    method: 'post',
    data
  })
}

// 删除replicaSet
export const deleteReplicaSetApi = (data) => {
  return request({
    url: '/k8s/replicaSet/deleteReplicaSet',
    method: 'post',
    data
  })
}
