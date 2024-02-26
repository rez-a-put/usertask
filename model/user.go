package model

import "time"

type User struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name" binding:"required,min=3,max=255"`
	Email      string     `json:"email" binding:"required,email,max=255"`
	Password   string     `json:"password,omitempty" binding:"required"`
	Status     string     `json:"status,omitempty"`
	DataStatus int8       `json:"-"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
