package handler

import (
	"github.com/sllt/kite-layout/api/v1"
	"github.com/sllt/kite-layout/internal/service"
	"github.com/sllt/kite/pkg/kite"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// Register godoc
// @Summary 用户注册
// @Schemes
// @Description 目前只支持邮箱登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.RegisterRequest true "params"
// @Success 200 {object} v1.Response
// @Router /register [post]
func (h *UserHandler) Register(ctx *kite.Context) (any, error) {
	req := new(v1.RegisterRequest)
	if err := ctx.Bind(req); err != nil {
		return nil, err
	}

	if err := h.userService.Register(ctx, req); err != nil {
		h.logger.Errorf("userService.Register error: %v", err)
		return nil, err
	}

	return nil, nil
}

// Login godoc
// @Summary 账号登录
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /login [post]
func (h *UserHandler) Login(ctx *kite.Context) (any, error) {
	var req v1.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		return nil, v1.ErrUnauthorized
	}

	return v1.LoginResponseData{
		AccessToken: token,
	}, nil
}

// GetProfile godoc
// @Summary 获取用户信息
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} v1.GetProfileResponse
// @Router /user [get]
func (h *UserHandler) GetProfile(ctx *kite.Context) (any, error) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		return nil, v1.ErrUnauthorized
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		return nil, v1.ErrBadRequest
	}

	return user, nil
}

// UpdateProfile godoc
// @Summary 修改用户信息
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body v1.UpdateProfileRequest true "params"
// @Success 200 {object} v1.Response
// @Router /user [put]
func (h *UserHandler) UpdateProfile(ctx *kite.Context) (any, error) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		return nil, v1.ErrUnauthorized
	}

	var req v1.UpdateProfileRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		return nil, v1.ErrInternalServerError
	}

	return nil, nil
}
