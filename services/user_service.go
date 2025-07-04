package services

import (
	"user/micro/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

// PingDataBase checks the connection to the MongoDB database.
// It returns true if the connection is successful, otherwise false.
func (us *UserService) PingDataBase() bool {
	return us.UserRepository.PingDataBase() == nil
}
