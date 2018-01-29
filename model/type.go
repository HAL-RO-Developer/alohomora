package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type User struct {
	Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
type Key struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Alive       bool   `json:"alive"`
}

type Token struct {
	Model
	UserID uint   `json:"user_id"`
	Body   string `json:"body"`
}
