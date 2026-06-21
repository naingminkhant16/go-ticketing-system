package entity

import (
	"time"
)

type UserRole string

const (
	SystemAdmin UserRole = "system-admin"
	SuperAdmin  UserRole = "super-admin"
	Admin       UserRole = "admin"
	Customer    UserRole = "customer"
)

type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
)

type User struct {
	Base
	Name       string     `gorm:"size:255;not null" json:"name"`
	Email      string     `gorm:"unique;not null" json:"email"`
	Password   string     `gorm:"not null;size:255" json:"-"`
	Role       UserRole   `gorm:"type:char(20);default:'customer'" json:"role"`
	Status     UserStatus `gorm:"type:char(20);default:'active'" json:"status"`
	VerifiedAt time.Time  `gorm:"default:null" json:"verified_at"`
	Gender     string     `gorm:"size:20;default:null" json:"gender"`
	Dob        time.Time  `gorm:"default:null" json:"dob"`

	// Relations
	Seats []Seat `gorm:"foreignKey:ReservedBy" json:"seats"`
}
