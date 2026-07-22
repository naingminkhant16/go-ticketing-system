package service

import (
	"testing"
	"ticketing-system/entity/dto"
)

func TestCreateReturnsErrorWhenCoverImageIsMissing(t *testing.T) {
	svc := &EventService{}

	input := &dto.EventCreateDto{
		Name:        "Launch Event",
		Description: "A test event",
		Location:    "Yangon",
		LocationURL: "https://example.com",
		StartDate:   "2026-07-21",
		EndDate:     "2026-07-21",
		AvailableAt: "2026-07-21",
		Organizer:   "Organizer",
		IssueCenter: "Center",
		EventDays:   []string{"2026-07-21"},
	}

	event, err := svc.Create(input)
	if err == nil {
		t.Fatal("expected an error when cover image is missing")
	}
	if event != nil {
		t.Fatal("expected no event to be created")
	}
}
