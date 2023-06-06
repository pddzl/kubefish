// pod

const statusPodMap: Record<string, podRes> = {
  'Running': 'success'
}

type podRes = 'success' | 'info'

export const statusPodFilter = (status: string): podRes => {
  return statusPodMap[status] || 'info'
}

// deploymentDetail -> replicaSet

type rsRes = 'success' | 'danger' | 'info'

const statusRsMap: Record<string, rsRes> = {
  'True': 'success',
  'False': 'danger',
  'Unknown': 'info',
}

export const statusRsFilter = (status: string): rsRes => {
  return statusRsMap[status] || 'info'
}
