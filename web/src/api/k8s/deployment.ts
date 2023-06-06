import { request } from "@/utils/service"
import { commonListReq, commonReq } from "./entry"

export interface DeploymentBrief {
  name: string
  Namespace: string
  AvailableReplicas: number
  Replicas: number
  CreationTimestamp: string
}

interface DeploymentBriefList extends PageInfo {
  list: DeploymentBrief[]
  total: number
}

// 获取deployment list
export const getDeploymentsApi = (data: commonListReq) => {
  return request<ApiResponseData<DeploymentBriefList>>({
    url: "/k8s/deployment/getDeployments",
    method: "post",
    data
  })
}

// 获取deployment详情
export const getDeploymentDetailApi = (data: commonReq) => {
  return request<ApiResponseData<any>>({
    url: "/k8s/deployment/getDeploymentDetail",
    method: "post",
    data
  })
}

// 删除deployment
export const deleteDeploymentApi = () => {}
