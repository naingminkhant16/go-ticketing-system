package service

import (
	"ticketing-system/entity"
	"ticketing-system/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) GetAll() ([]entity.User, error) {
	return us.userRepository.GetAll()
}

func (us *UserService) GetByEmail(email string) (*entity.User, error) {
	return us.userRepository.GetByEmail(email)
}

func (us *UserService) Create(user entity.User) (*entity.User, error) {
	return us.userRepository.Save(user)
}
