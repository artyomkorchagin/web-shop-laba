package types

type Category struct {
	ID        uint `json:"category_id"`
	Name      uint `json:"user_id"`
	Parent_ID uint `json:"parent_category_id"`
}

type CreateCategoryRequest struct {
	Name string `form:"name"`
}
