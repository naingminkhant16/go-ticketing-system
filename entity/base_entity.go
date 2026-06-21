package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey;" json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}
