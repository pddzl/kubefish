package request

type ResourceRaw struct {
	Name      string `json:"name" comment:"资源名称" validate:"required"`
	Resource  string `json:"resource" comment:"资源类型" validate:"required"`
	Namespace string `json:"namespace" comment:"命名空间"`
}
