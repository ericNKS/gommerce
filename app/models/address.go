package models

type Address struct {
	model
	Address string `json:"address" validate:"required"`
	Street  string `gorm:"index;not null" json:"street" validate:"required"`
	City    string `gorm:"index;not null" json:"city" validate:"required"`
	State   string `gorm:"index;not null" json:"state" validate:"required"`
	Cep     string `gorm:"index;not null" json:"cep" validate:"required"`
	Country string `gorm:"index;not null" json:"country" validate:"required"`
	Number  int16  `gorm:"index;not null" json:"number" validate:"required"`
}
