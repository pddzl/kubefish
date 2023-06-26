import { getResourceRawApi } from "@/api/k8s/resource"

// 查看资源编排
export const viewOrch = async (name: string, resource: string, namespace = "") => {
  let formatData = ""
  const result = await getResourceRawApi({ name: name, resource: resource, namespace: namespace })
  if (result.code === 0) {
    formatData = JSON.stringify(result.data)
  }

  return formatData
}
