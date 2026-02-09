// versions:
// 	kite-cli v0.1.0
// 	kite v0.1.0
// 	source: user.proto

package user

import (
	"github.com/sllt/kite-layout/internal/service"
	"github.com/sllt/kite-layout/internal/types"
	"github.com/sllt/kite/pkg/kite"
)

// Register the gRPC service in your app using the following code in your main.go:
//
// user.RegisterUserServiceServerWithKite(app, user.NewUserServiceKiteServer(userService))
//
// UserServiceKiteServer defines the gRPC server implementation.

type UserServiceKiteServer struct {
	health      *healthServer
	userService service.UserService
}

// NewUserServiceKiteServerWithService creates a new instance with service dependency
func NewUserServiceKiteServerWithService(userService service.UserService) *UserServiceKiteServer {
	return &UserServiceKiteServer{
		health:      getOrCreateHealthServer(),
		userService: userService,
	}
}

func (s *UserServiceKiteServer) Register(ctx *kite.Context) (any, error) {
	// 获取 protobuf 请求
	reqWrapper := ctx.Request.(*RegisterRequestWrapper)
	req := reqWrapper.RegisterRequest

	// pb → types 转换
	input := &types.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.userService.Register(ctx, input); err != nil {
		return nil, err
	}

	return &RegisterResponse{}, nil
}

func (s *UserServiceKiteServer) Login(ctx *kite.Context) (any, error) {
	reqWrapper := ctx.Request.(*LoginRequestWrapper)
	req := reqWrapper.LoginRequest

	// pb → types 转换
	input := &types.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := s.userService.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	// types → pb 转换
	return &LoginResponse{
		AccessToken: output.AccessToken,
	}, nil
}

func (s *UserServiceKiteServer) GetProfile(ctx *kite.Context) (any, error) {
	reqWrapper := ctx.Request.(*GetProfileRequestWrapper)
	req := reqWrapper.GetProfileRequest

	output, err := s.userService.GetProfile(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// types → pb 转换
	return &GetProfileResponse{
		UserId:   output.UserId,
		Nickname: output.Nickname,
	}, nil
}

func (s *UserServiceKiteServer) UpdateProfile(ctx *kite.Context) (any, error) {
	reqWrapper := ctx.Request.(*UpdateProfileRequestWrapper)
	req := reqWrapper.UpdateProfileRequest

	// pb → types 转换
	input := &types.UpdateProfileInput{
		Nickname: req.Nickname,
		Email:    req.Email,
	}

	if err := s.userService.UpdateProfile(ctx, req.UserId, input); err != nil {
		return nil, err
	}

	return &UpdateProfileResponse{}, nil
}
