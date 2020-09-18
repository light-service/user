package model

import "time"

type UserToken struct {
	ID        int       `db:"id"`
	Token     string    `db:"token"`
	IP        string    `db:"ip"`
	UserID    int       `db:"user_id"`
	ExpireAt  time.Time `db:"expire_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
