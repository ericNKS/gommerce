package models

import "github.com/google/uuid"

type ShoppingItem struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Shopping   Shopping  `gorm:"foreignKey:ShoppingId"`
	ShoppingId uuid.UUID `gorm:"index;not null"`

	Product   Product   `gorm:"foreignKey:ProductId"`
	ProductId uuid.UUID `gorm:"index;not null"`

	Quantity int `gorm:"not null" json:"quantity"`
}
