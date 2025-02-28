package models

import (
	"github.com/google/uuid"
)

type Shopping struct {
	model
	User   User      `gorm:"foreignKey:UserId"`
	UserId uuid.UUID `gorm:"index;not null" json:"userId" validate:"required"`

	Products []Product `gorm:"many2many:shopping_items;" json:"products"`

	Status int8 `gorm:"default:0" json:"status"`

	Address   Address   `gorm:"foreignKey:AddressId"`
	AddressId uuid.UUID `gorm:"index;not null" json:"addressId" validate:"required"`
}
