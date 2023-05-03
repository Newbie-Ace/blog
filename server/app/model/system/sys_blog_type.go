package system

import "github.com/noobHuKai/app/model/common"

type SysBlogType struct {
	common.Model

	Name string `json:"name" gorm:"unique;comment:类型名称"`
}

func (SysBlogType) TableName() string {
	return "sys_blog_type"
}
