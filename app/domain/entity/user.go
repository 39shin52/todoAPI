package entity

import "time"

type User struct {
	ID        string    `db:"user_id"`
	Password  string    `db:"password"`
	Token     string    `db:"token"`
	UserName  string    `db:"user_name"`
	Mail      string    `db:"mail"`
	Work      string    `db:"work"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
