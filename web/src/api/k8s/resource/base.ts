import { request } from "@/utils/service"

export function getResourceRawApi(data: { name: string; resource: string; namespace?: string }) {
  return request<ApiResponseData<any>>({
    url: "/k8s/resource/getResourceRaw",
    method: "post",
    data
  })
}
