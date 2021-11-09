package models

import (
	"fmt"
	"time"
)

type Booking struct {
	BaseModel
	Customer   string    `json:"customer"`
	StartDate  time.Time `json:"start_date" binding:"required" gorm:"not null"`
	EndDate    time.Time `json:"end_date" binding:"required" gorm:"not null"`
	Notes      string    `json:"notes"`
	RoomID     uint      `json:"room_id" gorm:"not null"`
	ServiceIDs []uint    `json:"service_ids" gorm:"-"`
	Services   []Service `json:"-" gorm:"many2many:booking_services"`
}

func (b *Booking) Validate() error {
	if b.EndDate.Before(b.StartDate) {
		return fmt.Errorf("end date cannot be before start date")
	}

	return nil
}
