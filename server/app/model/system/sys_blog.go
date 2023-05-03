package system

import (
	"github.com/noobHuKai/app/model/common"
)

type SysBlog struct {
	common.Model

	Title        string `json:"title" gorm:"comment:博客标题"`
	Content      string `json:"content" gorm:"comment:博客内容"`
	UserID       uint   `json:"user_id" gorm:"comment:创作者ID"`
	TypeID       uint   `json:"type_id" gorm:"comment:类型ID"`
	Status       uint   `json:"status"  gorm:"default:0;comment:博客状态"`
	ViewCount    uint   `json:"view_count"  gorm:"default:0;comment:浏览数量"`
	CommentCount uint   `json:"comment_count" gorm:"default:0;comment:评论数量"`
	LikeCount    uint   `json:"like_count" gorm:"default:0;comment:点赞数量"`
	CollectCount uint   `json:"collect_count" gorm:"default:0;comment:收藏数量"`

	CoverImage string `json:"cover_image" gorm:"comment:封面图片"`

	// 外键
	SysUser     SysUser     `json:"user" gorm:"foreignKey:UserID"`
	SysBlogType SysBlogType `json:"blog_type" gorm:"foreignKey:TypeID"`
}

func (SysBlog) TableName() string {
	return "sys_blog"
}

type BlogContent struct {
	Blog SysBlog  `json:"blog"`
	Tags []SysTag `json:"tags"`
}
