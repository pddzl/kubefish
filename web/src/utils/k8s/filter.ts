// pod

type podRes = "success" | "info"

const statusPodMap: Record<string, podRes> = {
  Running: "success"
}

export const statusPodFilter = (status: string): podRes => {
  return statusPodMap[status] || "info"
}

// deploymentDetail | replicaSet

type rsRes = "success" | "danger" | "info"

const statusRsMap: Record<string, rsRes> = {
  True: "success",
  False: "danger",
  Unknown: "info"
}

export const statusRsFilter = (status: string): rsRes => {
  return statusRsMap[status] || "info"
}

// namespace

type nsRes = "success"

const statusNsMap: Record<string, nsRes> = {
  Active: "success"
}

export const statusNsFilter = (status: string): nsRes => {
  return statusNsMap[status] || "info"
}

// node

type nodeRes = "success" | "danger"

const statusNodeTypeMap: Record<string, nodeRes> = {
  True: "success",
  Unknown: "danger"
}

export const statusNodeTypeFilter = (status: string): nodeRes => {
  return statusNodeTypeMap[status] || "info"
}

const nodeStatusMap: Record<string, string> = {
  True: "Ready",
  Unknown: "NotReady"
}

export const nodeStatusFilter = (status: string) => {
  return nodeStatusMap[status] || "Unknown"
}
