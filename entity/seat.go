package entity

import "github.com/google/uuid"

type SeatType string

const (
	Regular SeatType = "regular"
	VIP     SeatType = "vip"
	Premium SeatType = "premium"
)

type SeatStatus string

const (
	Available SeatStatus = "available"
	Reserved  SeatStatus = "reserved"
	Sold      SeatStatus = "sold"
	Blocked   SeatStatus = "blocked"
)

type Seat struct {
	Base
	EventDayId uuid.UUID  `json:"event_day_id"`
	SeatNumber string     `gorm:"type:char(100);not null" json:"seat_number"`
	Section    string     `gorm:"type:varchar(50);not null" json:"section"`
	RowNumber  string     `gorm:"type:varchar(50);default:null" json:"row_number"`
	SeatType   SeatType   `gorm:"type:char(20);default:null" json:"seat_type"`
	Price      float32    `gorm:"type:float;not null" json:"price"`
	Status     SeatStatus `gorm:"type:char(20);default:'available'" json:"status"`
	Version    uint       `gorm:"type:int;default:0" json:"version"`
	ReservedBy uuid.UUID  `gorm:"type:char(36);default:null" json:"reserved_by"`
	BookingID  uuid.UUID  `gorm:"type:char(36);default:null" json:"booking_id"`
}
