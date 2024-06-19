package service

import (
	"capstone-project/model"
	"capstone-project/repository"
)

type userService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

type UserService interface {
	GetUserTable() (*model.User, error)
	CreateUser(user model.User) error
	GetUserByUsername(user model.User) error
	GetUserByEmail(user model.User) error
	GetUserById(id int) (*model.User, error)
	UpdateUserByEmail(email string, user model.User) error
	DeleteUserById(id int) error
}

func NewUserService(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) *userService {
	return &userService{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (s *userService) GetUserTable() (*model.User, error) {
	users, err := s.userRepo.GetUserTable()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) CreateUser(user model.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserByUsername(user model.User) error {
	_, err := s.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUserByEmail(user model.User) error {
	_, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUserById(id int) (*model.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) UpdateUserByEmail(email string, user model.User) error {
	return s.userRepo.UpdateUserByEmail(email, user)
}

func (s *userService) DeleteUserById(id int) error {
	return s.userRepo.DeleteUserById(id)
}
