package modelPG

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `json:"user_id"`
	Token     string    `json:"token" gorm:"unique"`
	ExpiresAt time.Time `json:"expires_at"`
}
