package system

import "github.com/noobHuKai/app/model/common"

type SysBlogTag struct {
	common.Model

	TagID  uint `json:"tag_id"  gorm:"comment:标签ID"`
	BlogID uint `json:"blog_id" gorm:"comment:博客ID"`

	// 外键
	SysTag  SysTag  `json:"tag" gorm:"foreignKey:TagID"`
	SysBlog SysBlog `json:"blog" gorm:"foreignKey:BlogID"`
}

func (SysBlogTag) TableName() string {
	return "sys_blog_tag"
}
