package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sllt/kite-layout/internal/model"
	"github.com/sllt/kite-layout/internal/repository"
	"github.com/sllt/kite-layout/pkg/log"
	"github.com/sllt/kite/pkg/kite/datasource"
	kiteSQL "github.com/sllt/kite/pkg/kite/datasource/sql"
	"github.com/stretchr/testify/assert"
)

var logger *log.Logger

// testDB wraps *sql.DB to satisfy infra.DB interface for testing.
type testDB struct {
	*sql.DB
}

func (t *testDB) Begin() (*kiteSQL.Tx, error) {
	return nil, errors.New("not implemented in test")
}

func (t *testDB) Select(_ context.Context, _ any, _ string, _ ...any) {}

func (t *testDB) HealthCheck() *datasource.Health { return nil }

func (t *testDB) Dialect() string { return "mysql" }

func setupRepository(t *testing.T) (repository.UserRepository, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	db := &testDB{DB: mockDB}
	repo := repository.NewRepository(logger, db)
	userRepo := repository.NewUserRepository(repo)

	return userRepo, mock
}

func TestUserRepository_Create(t *testing.T) {
	userRepo, mock := setupRepository(t)

	ctx := context.Background()
	user := &model.User{
		UserId:   "123",
		Nickname: "Test",
		Password: "password",
		Email:    "test@example.com",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.UserId, user.Nickname, user.Password, user.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := userRepo.Create(ctx, user)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.Id)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Update(t *testing.T) {
	userRepo, mock := setupRepository(t)

	ctx := context.Background()
	user := &model.User{
		Id:       1,
		UserId:   "123",
		Nickname: "Test",
		Password: "password",
		Email:    "test@example.com",
	}

	mock.ExpectExec("UPDATE users SET").
		WithArgs(user.Nickname, user.Password, user.Email, sqlmock.AnyArg(), user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := userRepo.Update(ctx, user)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetById(t *testing.T) {
	userRepo, mock := setupRepository(t)

	ctx := context.Background()
	userId := "123"
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "user_id", "nickname", "password", "email", "created_at", "updated_at"}).
		AddRow(1, "123", "Test", "password", "test@example.com", now, now)
	mock.ExpectQuery("SELECT id, user_id, nickname, password, email, created_at, updated_at FROM users WHERE user_id").
		WithArgs(userId).
		WillReturnRows(rows)

	user, err := userRepo.GetByID(ctx, userId)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "123", user.UserId)
	assert.Equal(t, "Test", user.Nickname)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByEmail(t *testing.T) {
	userRepo, mock := setupRepository(t)

	ctx := context.Background()
	email := "test@example.com"
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "user_id", "nickname", "password", "email", "created_at", "updated_at"}).
		AddRow(1, "123", "Test", "password", "test@example.com", now, now)
	mock.ExpectQuery("SELECT id, user_id, nickname, password, email, created_at, updated_at FROM users WHERE email").
		WithArgs(email).
		WillReturnRows(rows)

	user, err := userRepo.GetByEmail(ctx, email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@example.com", user.Email)

	assert.NoError(t, mock.ExpectationsWereMet())
}
