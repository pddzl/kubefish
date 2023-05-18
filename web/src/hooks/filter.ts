type status = "True" | "Unknown"

const nodeStatusType: Record<status, string> = {
  True: "success",
  Unknown: "danger"
}

export const nodeStatusTypeFilter = (status: status): string => {
  return nodeStatusType[status] || "info"
}

const nodeStatus: Record<status, string> = {
  True: "Ready",
  Unknown: "NotReady"
}

export const nodeStatusFilter = (status: status) => {
  return nodeStatus[status] || "Unknown"
}

// Pod

type podStatus = "Running"

const podStatusMap: Record<podStatus, string> = {
  Running: "success"
}

export const PodStatusFilter = (status: podStatus) => {
  return podStatusMap[status] || "info"
}

// namespace

type namespaceStatus = "Active"

const namespaceStatusMap = {
  Active: "success"
}

export const NamespaceStatusFilter = (status: namespaceStatus) => {
  return namespaceStatusMap[status] || "info"
}
