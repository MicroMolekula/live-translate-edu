package services

import (
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/repository"
	"github.com/live-translate-edu/internal/utils"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) CreateUser(userDto *dto.UserCreateDTO) error {
	userModel := dto.UserCreateDTOToModel(userDto)
	hashPassword, err := utils.HashPassword(userModel.Password)
	if err != nil {
		return err
	}
	userModel.Password = hashPassword
	if err = us.userRepository.Create(userModel); err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserByEmail(email string) (*dto.UserFullDTO, error) {
	userModel, err := us.userRepository.FindOneByEmail(email)
	if err != nil {
		return nil, err
	}
	return dto.UserModelToFullDTO(userModel), nil
}
