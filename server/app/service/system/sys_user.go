package system

import (
	"errors"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/system"
	"gorm.io/gorm"
)

type UserService struct{}

// Register 用户注册
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(g.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, g.ErrUserAlreadyExist
	}
	err = g.DB.Create(&u).Error
	return u, err
}

// Login 用户登录
func (userService *UserService) Login(username, password string) (userInter *system.SysUser, err error) {
	var user system.SysUser

	err = g.DB.Where(&system.SysUser{Username: username}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, g.ErrUserNotFound
		}
		return nil, err
	}
	if user.Password != password {
		return nil, g.ErrUserPassWordWrong
	}
	return &user, err
}

// GetUserInfo 获取用户信息
func (userService *UserService) GetUserInfo(uid uint64) (userInter *system.SysUser, err error) {
	var user system.SysUser

	err = g.DB.First(&user, uid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, g.ErrUserNotFound
		}
		return nil, err
	}
	return &user, err
}
