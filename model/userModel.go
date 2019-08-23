package model

// UserModel 用户模型, 使用 struct tag
type UserModel struct {
	// struct tag 绑定多个 key:value 使用空格分隔，binding 进行表单验证
	// 更多验证规则: https://godoc.org/gopkg.in/go-playground/validator.v9
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}
