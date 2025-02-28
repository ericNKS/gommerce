package models

import "github.com/google/uuid"

type Product struct {
	model
	Name       string    `gorm:"index;not null" json:"name" validate:"required,min=4"`
	CategoryID uuid.UUID `json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
	Amount     int16     `json:"amount" validate:"required,min=0"`
	Price      int32     `json:"price" validate:"required,min=0"`
	Images     []string  `gorm:"serializer:json" json:"images" validate:"required"`
}
