package services

import (
	"errors"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	UserRepository     *repository.UserRepository
	UserRoleRepository *repository.UserRoleRepository
}

func (s *UserService) GetById(id string) (*model.User, error) {
	return s.UserRepository.GetById(id)
}

func (s *UserService) Login(email string, password string) (*model.User, error) {
	user, err := s.UserRepository.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) AlreadyRegistered(email string) bool {
	if _, err := s.UserRepository.GetByEmail(email); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *UserService) Register(input *model.Register) (*model.User, error) {
	role, err := s.UserRoleRepository.GetUserRole()

	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return s.UserRepository.Save(&model.User{
		ID:          uuid.Must(uuid.NewRandom()).String(),
		Name:        input.Name,
		Email:       input.Email,
		Password:    string(hash),
		DateOfBirth: input.DateOfBirth,
		Gender:      input.Gender,
		Address:     input.Address,
		UserRoleID:  role.ID,
		UserRole:    role,
	})
}

func (s *UserService) UpdateAccount(input *model.Update) (*model.User, error) {
	user, err := s.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Gender = input.Gender
	user.DateOfBirth = input.DateOfBirth
	user.Email = input.Email
	user.Address = input.Address

	return s.UserRepository.Update(user)
}

func (s *UserService) DeleteAccount(input *model.DeleteAccount) (*model.User, error) {
	user, err := s.GetById(input.ID)
	if err != nil {
		return nil, err
	}
	return s.UserRepository.Delete(user)
}
