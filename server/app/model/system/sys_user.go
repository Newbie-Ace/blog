package system

import (
	"github.com/noobHuKai/app/model/common"
)

type SysUser struct {
	common.Model

	Username  string `json:"username" gorm:"index;unique;comment:用户登录名"`                                            // 用户登录名
	Password  string `json:"-"  gorm:"comment:用户登录密码"`                                                              // 用户登录密码
	NickName  string `json:"nickname" gorm:"default:user;comment:用户昵称"`                                             // 用户昵称
	HeaderImg string `json:"header_img" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Phone     string `json:"phone"  gorm:"comment:用户手机号"`                                                           // 用户手机号
	Email     string `json:"email"  gorm:"comment:用户邮箱"`                                                            // 用户邮箱
	Age       int    `json:"age"  gorm:"comment:用户年龄"`                                                              //用户年龄
	Status    int    `json:"status"  gorm:"default:0;comment:用户状态"`                                                 //用户状态
	Role      string `json:"role"  gorm:"default:user;comment:用户角色"`                                                //用户角色
}

func (SysUser) TableName() string {
	return "sys_user"
}
