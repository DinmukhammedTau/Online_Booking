package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);unique;not null"`
}
