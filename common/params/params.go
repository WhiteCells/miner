package params

type PageParams struct {
	PageNum  int `form:"page_num" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1"`
}
