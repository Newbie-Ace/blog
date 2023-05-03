package request

import (
	"github.com/noobHuKai/app/model/common/request"
)

// CreateBlogType structure
type CreateBlogType struct {
	Name string `json:"name"` // 类型名称
}

// CreateBlogTag structure
type CreateBlogTag struct {
	Name string `json:"name"` // tag 名称
}

// CreateBlog structure
type CreateBlog struct {
	Title   string `json:"title"`   // 标题
	Content string `json:"content"` // 内容
	Tags    []uint `json:"tags"`    // 标签
	TypeID  uint   `json:"type_id"` // 博客类型
}

// ListBlog structure
type ListBlog struct {
	request.Page
}
