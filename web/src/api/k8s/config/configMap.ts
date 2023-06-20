import { request } from "@/utils/service"
import { type commonListReq, type commonReq } from "@/api/k8s/entry"

// 获取configMap list
export interface ConfigMapBrief {
  name: string
  namespace: string
  creationTimestamp: string
}

interface ConfigMapBriefList extends PageInfo {
  list: ConfigMapBrief[]
  total: number
}

export const getConfigMapListApi = (data: commonListReq) => {
  return request<ApiResponseData<ConfigMapBriefList>>({
    url: "/k8s/configMap/getConfigMapList",
    method: "post",
    data
  })
}

// 获取configMap详情

// export interface ConfigMapDetail {
//   metadata: {}
//   spec: {
//     selector: object
//     updateStrategy: { type: string; rollingUpdate: { maxUnavailable: number; maxSurge: number } }
//   }
//   status: {
//     currentNumberScheduled: number
//     numberMisscheduled: number
//     desiredNumberScheduled: number
//     numberReady: number
//     updatedNumberScheduled: number
//     numberAvailable: number
//     conditions: []
//   }
// }

// export const getConfigMapDetailApi = (data: commonReq) => {
//   return request<ApiResponseData<ConfigMapDetail>>({
//     url: "/k8s/configMap/getConfigMapDetail",
//     method: "post",
//     data
//   })
// }

// 删除configMap
export const deleteConfigMapApi = (data: commonReq) => {
  return request<ApiResponseData<null>>({
    url: "/k8s/configMap/deleteConfigMap",
    method: "post",
    data
  })
}
