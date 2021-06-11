package models

import (
	"time"
)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
}

type Users struct {
	Users []User `json:"users"`
}
