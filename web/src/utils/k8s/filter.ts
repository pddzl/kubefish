// pod

const statusPodMap: Record<string, podRes> = {
  'Running': 'success'
}

type podRes = 'success' | 'info'

export const StatusPodFilter = (status: string): podRes => {
  return statusPodMap[status] || 'info'
}

export const StatusRsFilter = () => {}
