package repository

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func (repo *BaseRepository[T]) GetAll() ([]T, error) {
	var data []T
	result := repo.db.Last(&data)

	if result.Error != nil {
		log.Printf("Error retrieving records: %v", result.Error)
		return nil, result.Error
	}

	return data, nil
}

func (repo *BaseRepository[T]) GetById(id uuid.UUID) (*T, error) {
	var data T
	result := repo.db.First(&data, "id = ?", id.String())
	if result.Error != nil {
		log.Printf("Error retrieving records: %v", result.Error)
		return nil, result.Error
	}
	return &data, nil
}

func (repo *BaseRepository[T]) Save(data T) (*T, error) {
	result := repo.db.Save(&data)
	if result.Error != nil {
		log.Printf("Error saving records: %v", result.Error)
		return nil, result.Error
	}
	return &data, nil
}
