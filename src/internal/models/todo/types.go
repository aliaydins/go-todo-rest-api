package todo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Todo      string    `gorm:"size:255;not null;"`
	Done      bool      `gorm:"default:false;"`
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
