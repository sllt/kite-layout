package types

// RegisterInput 用户注册输入
type RegisterInput struct {
	Email    string
	Password string
}

// LoginInput 用户登录输入
type LoginInput struct {
	Email    string
	Password string
}

// LoginOutput 登录输出
type LoginOutput struct {
	AccessToken string
}

// UserOutput 用户信息输出
type UserOutput struct {
	UserId   string
	Nickname string
}

// UpdateProfileInput 更新用户资料输入
type UpdateProfileInput struct {
	Nickname string
	Email    string
}
