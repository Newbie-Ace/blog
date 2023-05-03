package system

import "github.com/noobHuKai/app/model/common"

type SysTag struct {
	common.Model

	Name string `json:"name" gorm:"unique;comment:标签名称"`
}

func (SysTag) TableName() string {
	return "sys_tag"
}
