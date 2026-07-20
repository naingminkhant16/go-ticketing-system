package repository

import (
	"ticketing-system/entity"

	"gorm.io/gorm"
)

type EventRepository struct {
	BaseRepository[entity.Event]
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		BaseRepository: BaseRepository[entity.Event]{db: db},
	}
}
