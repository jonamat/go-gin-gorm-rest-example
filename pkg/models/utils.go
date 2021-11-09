package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
