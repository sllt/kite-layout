package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	v1 "github.com/sllt/kite-layout/api/v1"
	"github.com/sllt/kite-layout/internal/model"
)

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
	user.UpdatedAt = time.Now()

	q := r.GetQuerier(ctx)
	_, err := q.ExecContext(ctx,
		"UPDATE users SET nickname = ?, password = ?, email = ?, updated_at = ? WHERE id = ?",
		user.Nickname, user.Password, user.Email, user.UpdatedAt, user.Id,
	)
	return err
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.User, error) {
	q := r.GetQuerier(ctx)
	var user model.User
	err := q.QueryRowContext(ctx,
		"SELECT id, user_id, nickname, password, email, created_at, updated_at FROM users WHERE user_id = ?",
		userId,
	).Scan(&user.Id, &user.UserId, &user.Nickname, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	q := r.GetQuerier(ctx)
	var user model.User
	err := q.QueryRowContext(ctx,
		"SELECT id, user_id, nickname, password, email, created_at, updated_at FROM users WHERE email = ?",
		email,
	).Scan(&user.Id, &user.UserId, &user.Nickname, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
