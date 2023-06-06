import { request } from "@/utils/service"
import { reqNamespace } from "./entry"

export interface replicaSetBrief {
  name: string
  namespace: string
  availableReplicas: number
  replicas: number
  creationTimestamp: string
}

export interface ReplicaSetBriefList extends PageInfo {
  list: replicaSetBrief[]
  total: number
}

// 获取replicaSet list
export const getReplicaSetsApi = (data: reqNamespace) => {
  return request<ApiResponseData<ReplicaSetBriefList>>({
    url: '/k8s/replicaSet/getReplicaSets',
    method: 'post',
    data
  })
}

interface rsReq {
  namespace: string
  replicaSet: string
}

// 删除replicaSet
export const deleteReplicaSetApi = (data: rsReq) => {
  return request<ApiResponseData<null>>({
    url: '/k8s/replicaSet/deleteReplicaSet',
    method: 'post',
    data
  })
}

export interface ReplicaSetDetail {
  metadata: Object,
  spec: {
    replicas: number,
    selector: {}
  },
  status: {
    replicas: number,
    fullyLabeledReplicas: number,
    readyReplicas: number,
    availableReplicas: number,
    conditions: []
  }
}

// 获取replicaSet详情
export const getReplicaSetDetailApi = (data: rsReq) => {
  return request<ApiResponseData<ReplicaSetDetail>>({
    url: '/k8s/replicaSet/getReplicaSetDetail',
    method: 'post',
    data
  })
}

interface rsPodReq extends rsReq, PageInfo {}

// 获取replicaSet关联的pod
export const getReplicaSetPodsApi = (data: rsPodReq) => {
  return request<ApiResponseData<any>>({
    url: '/k8s/replicaSet/getReplicaSetPods',
    method: 'post',
    data
  })
}

// 获取replicaSet关联的service
export const getReplicaSetServicesApi = (data: rsReq) => {
  return request<ApiResponseData<any>>({
    url: '/k8s/replicaSet/getReplicaSetServices',
    method: 'post',
    data
  })
}
