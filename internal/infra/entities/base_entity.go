package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID         uuid.UUID  `json:"id" gorm:"primaryKey"`
	MacAddress string     `json:"mac_address" gorm:"column:mac_address;not null"`
	Timestamp  time.Time  `json:"timestamp" gorm:"column:timestamp;not null"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at;not null"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (b *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return
}
