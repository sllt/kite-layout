package v1

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" binding:"required" label:"密码"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" binding:"required" label:"密码"`
}
type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" label:"昵称"`
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
}
type GetProfileResponseData struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}
