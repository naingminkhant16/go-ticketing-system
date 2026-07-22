package service

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	apperror "ticketing-system/common/error"
	"ticketing-system/entity"
	"ticketing-system/entity/dto"
	"ticketing-system/repository"
	"ticketing-system/service/storage"
	"time"
)

type EventService struct {
	eventRepository *repository.EventRepository
	s3Service       *storage.S3
}

func NewEventService(eventRepository *repository.EventRepository, s3 *storage.S3) *EventService {
	return &EventService{eventRepository: eventRepository, s3Service: s3}
}

func (es *EventService) GetAllEvents() ([]entity.Event, error) {
	return es.eventRepository.GetAll()
}

func (es *EventService) Create(dto *dto.EventCreateDto) (*entity.Event, error) {
	if dto == nil {
		return nil, apperror.BadRequest("event payload is required")
	}
	if dto.CoverImage == nil {
		return nil, apperror.BadRequest("cover image is required")
	}

	startDate, _ := time.Parse("2006-01-02", dto.StartDate)
	endDate, _ := time.Parse("2006-01-02", dto.EndDate)
	availableAt, _ := time.Parse("2006-01-02", dto.AvailableAt)

	// TODO : seats creations for each event day
	var eventDays []entity.EventDay
	for _, e := range dto.EventDays {
		date, _ := time.Parse("2006-01-02", e)
		eventDays = append(eventDays, entity.EventDay{
			Date: &date,
		})
	}

	key := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(dto.CoverImage.Filename))

	file, err := dto.CoverImage.Open()
	if err != nil {
		log.Println(err)
		return nil, apperror.BadRequest("failed to open cover image")
	}
	defer file.Close()

	// upload cover photo file to s3
	key, err = es.s3Service.Upload(context.Background(), key, file)

	if err != nil {
		log.Println("Failed to upload file to S3 : ", err)
		return nil, err
	}

	event := entity.Event{
		Name:           dto.Name,
		Description:    dto.Description,
		Location:       dto.Location,
		LocationURL:    dto.LocationURL,
		StartDate:      &startDate,
		EndDate:        &endDate,
		AvailableAt:    &availableAt,
		Organizer:      dto.Organizer,
		CoverImage:     key,
		IssueCenter:    dto.IssueCenter,
		TotalSeats:     0,
		AvailableSeats: 0,
		EventDays:      eventDays,
	}

	e, err := es.eventRepository.Save(event)
	if err != nil {
		log.Println("Failed to create event", e)
		return nil, err
	}

	return e, nil
}
