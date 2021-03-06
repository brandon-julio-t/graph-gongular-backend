package services

import (
	"errors"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/brandon-julio-t/graph-gongular-backend/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserService struct {
	UserRepository     *repository.UserRepository
	UserRoleRepository *repository.UserRoleRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		UserRepository:     &repository.UserRepository{DB: db},
		UserRoleRepository: &repository.UserRoleRepository{DB: db},
	}
}

func (s *UserService) GetAllExcept(user *model.User) ([]*model.User, error) {
	return s.UserRepository.GetAllExcept(user)
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

	return s.UserRepository.Save(
		&model.User{
			ID:          uuid.Must(uuid.NewRandom()).String(),
			Name:        input.Name,
			Email:       input.Email,
			Password:    string(hash),
			DateOfBirth: input.DateOfBirth,
			Gender:      input.Gender,
			Address:     input.Address,
			UserRoleID:  role.ID,
			UserRole:    role,
		},
	)
}

func (s *UserService) UpdateAccount(id string, input *model.UpdateUser) (*model.User, error) {
	user, err := s.GetById(id)

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

func (s *UserService) DeleteAccount(id string) (*model.User, error) {
	user, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	return s.UserRepository.Delete(user)
}
