package model

import (
	"time"
)

type User struct {
	ID          int       `db:"id"`
	Username    string    `db:"username"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Phone       string    `db:"phone"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	LastLoginAt time.Time `db:"last_login_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
}
