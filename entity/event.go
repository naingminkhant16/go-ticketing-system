package entity

import (
	"time"

	"github.com/google/uuid"
)

type EventStatus string

const (
	Upcoming  EventStatus = "upcoming"
	OnSale    EventStatus = "on_sale"
	SoldOut   EventStatus = "sold_out"
	Cancelled EventStatus = "cancelled"
)

type Event struct {
	Base
	Name           string      `gorm:"type:varchar(255)" json:"name"`
	Description    string      `gorm:"type:varchar(512);default:null" json:"description"`
	Location       string      `gorm:"type:varchar(512);not null" json:"location"`
	LocationURL    string      `gorm:"type:varchar(255);default:null" json:"location_url"`
	StartDate      time.Time   `gorm:"type:date;not null" json:"start_date"`
	EndDate        time.Time   `gorm:"type:date;not null" json:"end_date"`
	AvailableAt    time.Time   `gorm:"type:timestamp;not null" json:"available_at"`
	CoverImage     string      `gorm:"type:varchar(255);not null" json:"cover_image"`
	Organizer      string      `gorm:"type:varchar(255);not null" json:"organizer"`
	IsPublished    bool        `gorm:"type:boolean;default:false" json:"is_published"`
	IssueCenter    string      `gorm:"type:varchar(255);not null" json:"issue_center"`
	TotalSeats     int         `gorm:"type:int;not null" json:"total_seats"`
	AvailableSeats int         `gorm:"type:int;not null" json:"available_seats"`
	Status         EventStatus `gorm:"type:char(20);not null;default:upcoming" json:"status"`

	// Relations
	EventDays []EventDay `json:"event_days"`
}

type EventDay struct {
	Base
	Date     time.Time `gorm:"type:date;not null" json:"date"`
	IsActive bool      `gorm:"type:boolean;default:false" json:"is_active"`
	EventID  uuid.UUID `json:"event_id"`

	// Relations
	Seats []Seat `json:"seats"`
}
