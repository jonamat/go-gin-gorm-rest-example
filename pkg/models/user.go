package models

import (
	"fmt"
	"net/mail"
)

type User struct {
	BaseModel
	Email    string `json:"email" gorm:"not null;unique;index"`
	Password string `json:"password" gorm:"not null"`
	Level    string `json:"level" gorm:"not null"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	if u.Level != "" {
		return fmt.Errorf("you cannot change your permission level")
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return fmt.Errorf("email is not valid")
	}

	if len(u.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	return nil
}
