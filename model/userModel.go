package model

import (
	"go-web/initdb"
	"log"
)

// UserModel 用户模型, 使用 struct tag
type UserModel struct {
	// struct tag 绑定多个 key:value 使用空格分隔，binding 进行表单验证
	// 更多验证规则: https://godoc.org/gopkg.in/go-playground/validator.v9
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}

// Save 保存用户
func (user *UserModel) Save() int64 {
	result, e := initdb.Db.Exec(
		"insert into golang.user (email, password) values (?, ?)",
		user.Email,
		user.Password,
	)

	if e != nil {
		log.Panicln("user insert error", e.Error())
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}

	return id
}
