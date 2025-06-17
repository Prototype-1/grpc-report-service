package model

import (
    "github.com/google/uuid"
    "time"
)

type Report struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    UserID    string    `gorm:"not null"`
    CreatedAt time.Time
}
