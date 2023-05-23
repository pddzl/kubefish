import { request } from "@/utils/service"

interface namespacePageInfo extends PageInfo {
  list: any[]
  total: number
}

export function getNamespacesApi(data: PageInfo) {
  return request<ApiResponseData<namespacePageInfo>>({
    url: "/k8s/namespace/getNamespaces",
    method: "post",
    data
  })
}

// 获取某个命名空间详情
export function getNamespaceDetailApi(params: { name: string }) {
  return request<ApiResponseData<any>>({
    url: "/k8s/namespace/getNamespaceDetail",
    method: "get",
    params
  })
}

// 删除命名空间
export function deleteNamespaceApi(data: { name: string }) {
  return request<ApiResponseData<null>>({
    url: "/k8s/namespace/deleteNamespace",
    method: "post",
    data
  })
}

// 获取全部命名空间名称
export function getNamespaceNameApi() {
  return request<ApiResponseData<string[]>>({
    url: "/k8s/namespace/getNamespaceName",
    method: "get"
  })
}
