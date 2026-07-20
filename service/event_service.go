package service

import (
	"log"
	"ticketing-system/entity"
	"ticketing-system/entity/dto"
	"ticketing-system/repository"
	"time"
)

type EventService struct {
	eventRepository *repository.EventRepository
}

func NewEventService(eventRepository *repository.EventRepository) *EventService {
	return &EventService{eventRepository: eventRepository}
}

func (es *EventService) GetAllEvents() ([]entity.Event, error) {
	return es.eventRepository.GetAll()
}

func (es *EventService) Create(dto *dto.EventCreateDto) (*entity.Event, error) {

	startDate, _ := time.Parse("2006-01-02", dto.StartDate)
	endDate, _ := time.Parse("2006-01-02", dto.EndDate)
	availableAt, _ := time.Parse("2006-01-02", dto.AvailableAt)

	// TODO : upload cover photo file
	// TODO : seats creations for each event day
	var eventDays []entity.EventDay
	for _, e := range dto.EventDays {
		date, _ := time.Parse("2006-01-02", e)
		eventDays = append(eventDays, entity.EventDay{
			Date: &date,
		})
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
		CoverImage:     "filePath", // TODO : replace with real file path
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
