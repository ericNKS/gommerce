package models

import "github.com/google/uuid"

type User struct {
	Model
	Name     string
	Email    string `gorm:"index:idx_email,unique"`
	Password string
	Rule     int8 `gorm:"default:0"`
}

type UserCreate struct {
	Name     string `json:"name" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Rule     int8   `json:"rule"`
}

type UserResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Rule int8      `json:"rule"`
}
