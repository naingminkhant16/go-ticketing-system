package dto

import "mime/multipart"

type EventCreateDto struct {
	Name        string                `form:"name" binding:"required,min=3,max=50"`
	Description string                `form:"description" binding:"required,min=3,max=150"`
	Location    string                `form:"location" binding:"required,min=3,max=100"`
	LocationURL string                `form:"location_url" binding:"required,min=3,max=150"`
	StartDate   string                `form:"start_date"  binding:"required,datetime=2006-01-02"`
	EndDate     string                `form:"end_date"  binding:"required,datetime=2006-01-02"`
	AvailableAt string                `form:"available_at" binding:"required,datetime=2006-01-02"`
	CoverImage  *multipart.FileHeader `form:"cover_image" binding:"required"`
	Organizer   string                `form:"organizer" binding:"required,min=3,max=100"`
	IssueCenter string                `form:"issue_center" binding:"required,min=3,max=150"`
	EventDays   []string              `form:"event_days" binding:"required,dive,datetime=2006-01-02"`
}
