package repository

import (
	"ticketing-system/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: BaseRepository[entity.User]{db: db},
	}
}

func (repo *UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := repo.db.Where(&entity.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	
	return &user, nil
}
