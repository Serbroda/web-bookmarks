package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	id, err := shortid.Generate()
	b.ID = id
	return err
}
