package models

import "fmt"

type Room struct {
	BaseModel
	Name      string    `json:"name" binding:"required" gorm:"not null"`
	MaxGuests uint8     `json:"max_guests" binding:"required" gorm:"not null"`
	Price     float32   `json:"price"`
	Notes     string    `json:"notes"`
	Bookings  []Booking `json:"-"`
}

func (r *Room) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	if r.MaxGuests < 1 {
		return fmt.Errorf("max_guests must be greater than 0")
	}
	if r.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}

	return nil
}
