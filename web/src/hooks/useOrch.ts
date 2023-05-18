import { getResourceRawApi } from "@/api/k8s/resource"

export const useOrch = () => {
  let formatData: string
  const viewOrch = async (name: string, resource: string, namespace = ""): Promise<string> => {
    const result = await getResourceRawApi({ name: name, resource: resource, namespace: namespace })
    if (result.code === 0) {
      formatData = JSON.stringify(result.data)
    }
    return formatData
  }

  return {
    viewOrch
  }
}
