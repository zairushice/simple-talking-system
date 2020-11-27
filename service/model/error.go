package model

import "errors"

var (
	ErrorNoSuchUserId  = errors.New("用户id不存在")
	ErrorWrongPassword = errors.New("输入的密码错误")
	ErrorExistUserId   = errors.New("已经存在该用户id")
)
