import { ref, Ref } from "vue"
import { getResourceRawApi } from "@/api/k8s/resource"

export const useOrch = (name: string, resource: string, namespace = "") => {
  const formatData = ref("")
  const viewOrch = async (): Promise<Ref<string>> => {
    const result = await getResourceRawApi({ name: name, resource: resource, namespace: namespace })
    if (result.code === 0) {
      formatData.value = JSON.stringify(result.data)
    }
    return formatData
  }

  return {
    viewOrch
  }
}
