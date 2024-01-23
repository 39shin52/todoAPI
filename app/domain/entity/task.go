package entity

import "time"

type Task struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	IsComplete  int       `db:"is_complete"`
	Created_at  time.Time `db:"created_at"`
	Updated_at  time.Time `db:"updated_at"`
}
