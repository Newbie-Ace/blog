package request

type Page struct {
	PageSize int `json:"page_size" form:"page_size"` // 页大小
	PageNum  int `json:"page_num" form:"page_num"`   // 页数
}

type ID struct {
	ID uint `uri:"name" binding:"required"` // id
}
