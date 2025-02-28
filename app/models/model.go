package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type model struct {
	ID        uuid.UUID      `gorm:"index:idx_id,primarykey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
