package resource

type DynamicResource struct {
	Content string `json:"content" validated:"required"`
}
