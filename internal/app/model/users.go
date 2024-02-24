package model

import "github.com/google/uuid"

type Users struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	FullName string
	Email    string
	UserName string
	Password string
	Role     string
}
