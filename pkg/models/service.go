package models

import "fmt"

type Service struct {
	BaseModel
	Name  string  `json:"name" binding:"required" gorm:"not null"`
	Price float32 `json:"price"`
}

func (s *Service) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if s.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}

	return nil
}
