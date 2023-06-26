import { request } from "@/utils/service"
import { commonListReq, type commonRelatedReq, type commonReq } from "../entry"

// 获取daemonSet list
export interface DaemonSetBrief {
  name: string
  namespace: string
  pods: string
  creationTimestamp: string
}

interface DaemonSetBriefList extends PageInfo {
  list: DaemonSetBrief[]
  total: number
}

export const getDaemonSetsApi = (data: commonListReq) => {
  return request<ApiResponseData<DaemonSetBriefList>>({
    url: "/k8s/daemonSet/getDaemonSets",
    method: "post",
    data
  })
}

// 获取daemonSet详情

export interface DaemonSetDetail {
  metadata: {}
  spec: {
    selector: object
    updateStrategy: { type: string; rollingUpdate: { maxUnavailable: number; maxSurge: number } }
  }
  status: {
    currentNumberScheduled: number
    numberMisscheduled: number
    desiredNumberScheduled: number
    numberReady: number
    updatedNumberScheduled: number
    numberAvailable: number
    conditions: []
  }
}

export const getDaemonSetDetailApi = (data: commonReq) => {
  return request<ApiResponseData<DaemonSetDetail>>({
    url: "/k8s/daemonSet/getDaemonSetDetail",
    method: "post",
    data
  })
}

// 删除daemonSet
export const deleteDaemonSetApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/daemonSet/deleteDaemonSet",
    method: "post",
    data
  })
}

// 获取daemonSet关联的pod

export const getDaemonSetPodsApi = (data: commonRelatedReq) => {
  return request<ApiResponseData<any>>({
    url: "/k8s/daemonSet/getDaemonSetPods",
    method: "post",
    data
  })
}
