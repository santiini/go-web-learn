package model

// UserModel 用户模型, 使用 struct tag
type UserModel struct {
	Email         string `form:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again"`
}
