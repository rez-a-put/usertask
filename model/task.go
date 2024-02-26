package model

import "time"

type Task struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	StatusInt   int8       `json:"-"`
	Status      string     `json:"status,omitempty"`
	DataStatus  int8       `json:"-"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
