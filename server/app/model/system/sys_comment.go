package system

import "github.com/noobHuKai/app/model/common"

type SysComment struct {
	common.Model

	Content string `json:"content" gorm:"comment:评论内容"`
	UserID  uint   `json:"user_id" gorm:"comment:评论人ID"`
	BlogID  uint   `json:"blog_id" gorm:"comment:博客ID"`
	ReplyID uint   `json:"reply_id" gorm:"comment:评论回复ID"`

	// 外键
	SysUser SysUser `json:"user" gorm:"foreignKey:UserID"`
	SysBlog SysBlog `json:"blog" gorm:"foreignKey:BlogID"`
}

func (SysComment) TableName() string {
	return "sys_comment"
}
