package model

import "time"

type User struct {
	Id        uint      `db:"id"`
	UserId    string    `db:"user_id"`
	Nickname  string    `db:"nickname"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
