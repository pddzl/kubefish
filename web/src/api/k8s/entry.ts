export interface commonListReq extends PageInfo {
  namespace: string
}

export interface commonReq {
  namespace: string
  name: string
}

export interface commonRelatedReq extends commonReq, PageInfo {}

export interface nameReq {
  name: string
}
