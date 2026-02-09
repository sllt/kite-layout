package service

import (
	"context"
	"github.com/sllt/kite-layout/internal/model"
	"github.com/sllt/kite-layout/internal/repository"
	"github.com/sllt/kite-layout/internal/types"
	"github.com/sllt/kite-layout/pkg/errcode"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	Register(ctx context.Context, input *types.RegisterInput) error
	Login(ctx context.Context, input *types.LoginInput) (*types.LoginOutput, error)
	GetProfile(ctx context.Context, userId string) (*types.UserOutput, error)
	UpdateProfile(ctx context.Context, userId string, input *types.UpdateProfileInput) error
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) Register(ctx context.Context, input *types.RegisterInput) error {
	// check username
	user, err := s.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return errcode.ErrInternalServerError
	}
	if err == nil && user != nil {
		return errcode.ErrEmailAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}
	user = &model.User{
		UserId:   userId,
		Email:    input.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		// TODO: other repo
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, input *types.LoginInput) (*types.LoginOutput, error) {
	user, err := s.userRepo.GetByEmail(ctx, input.Email)
	if err != nil || user == nil {
		return nil, errcode.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errcode.ErrUnauthorized
	}
	token, err := s.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return nil, err
	}

	return &types.LoginOutput{AccessToken: token}, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*types.UserOutput, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &types.UserOutput{
		UserId:   user.UserId,
		Nickname: user.Nickname,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, input *types.UpdateProfileInput) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	user.Email = input.Email
	user.Nickname = input.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
