package models

import "github.com/google/uuid"

type User struct {
	model
	Name     string
	Email    string `gorm:"index:idx_email,unique"`
	Password string
	Rule     []string `gorm:"serializer:json;default:'{\"user\"}'"`
}

type UserCreate struct {
	Name     string   `json:"name" validate:"required,min=4,max=100" msg:"Name is required and should have at least 4 characters"`
	Email    string   `json:"email" validate:"required,email" msg:"A valid email is required"`
	Password string   `json:"password" validate:"required,min=8" msg:"Password is required and should have at least 8 characters"`
	Rule     []string `json:"rule"`
}

type UserResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Rule []string  `json:"rule"`
}
