package models

import "gorm.io/gorm"

type DummyUser struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
