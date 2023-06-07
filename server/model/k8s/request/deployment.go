package request

// DmScale 伸缩
type DmScale struct {
	CommonReq
	Num uint `json:"num" validate:"required"`
}
