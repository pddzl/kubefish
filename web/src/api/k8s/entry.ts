
export interface reqNamespace extends PageInfo {
  namespace: string
}

export interface commonListReq extends PageInfo {
  namespace: string
}

export interface commonReq {
  namespace: string
  name: string
}

export interface commonRelatedReq extends commonReq, PageInfo {}