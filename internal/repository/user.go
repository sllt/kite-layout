package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/sllt/kite-layout/internal/model"
	"github.com/sllt/kite-layout/pkg/errcode"
)

var errNilUser = errors.New("user is nil")

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

func NewUserRepository(
	r *Repository,
) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if user == nil {
		return errNilUser
	}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	q := r.GetQuerier(ctx)
	result, err := q.ExecContext(ctx,
		"INSERT INTO users (user_id, nickname, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		user.UserId, user.Nickname, user.Password, user.Email, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err == nil {
		user.Id = uint(id)
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if user == nil {
		return errNilUser
	}

	user.UpdatedAt = time.Now()

	q := r.GetQuerier(ctx)
	_, err := q.ExecContext(ctx,
		"UPDATE users SET nickname = ?, password = ?, email = ?, updated_at = ? WHERE id = ?",
		user.Nickname, user.Password, user.Email, user.UpdatedAt, user.Id,
	)
	return err
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.User, error) {
	return r.queryOne(ctx, "SELECT * FROM users WHERE user_id = ?", userId, errcode.ErrNotFound)
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return r.queryOne(ctx, "SELECT * FROM users WHERE email = ?", email, nil)
}

func (r *userRepository) queryOne(ctx context.Context, query string, arg any, notFoundErr error) (*model.User, error) {
	q := r.GetQuerier(ctx)
	var user model.User
	err := q.Select(ctx, &user, query, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, notFoundErr
		}
		return nil, err
	}

	return &user, nil
}
