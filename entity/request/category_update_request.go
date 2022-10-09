package request

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required" json:"name"`
}
