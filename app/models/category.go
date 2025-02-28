package models

type Category struct {
	model
	Name string `gorm:"uniqueIndex" json:"name" validate:"required,min=4"`
}
