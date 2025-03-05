package services

import (
	"errors"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/repository"
	"github.com/live-translate-edu/internal/utils"
	"time"
)

var (
	ErrorInvalidCredentials = errors.New("invalid credentials")
)

type AuthService struct {
	jwtService  *JWTService
	repository  repository.IRepository
	userService *UserService
}

func NewAuthService(
	jwtService *JWTService,
	userRepository *repository.UserRepository,
	userService *UserService) *AuthService {
	return &AuthService{
		jwtService:  jwtService,
		repository:  userRepository,
		userService: userService,
	}
}

func (as *AuthService) Auth(authData dto.AuthDto) (string, error) {
	user, err := as.userService.GetUserByEmail(authData.Login)
	if err != nil {
		return "", ErrorInvalidCredentials
	}
	if err = utils.CheckPassword(user.Password, authData.Password); err != nil {
		return "", ErrorInvalidCredentials
	}
	token, err := as.jwtService.GenerateTokenByUser(user, 60*24)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as *AuthService) Me(tokenString string) (string, error) {
	data, err := as.jwtService.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return data.Username, nil
}

func (as *AuthService) VerifyToken(tokenString string) (*dto.UserDTO, error) {
	data, err := as.jwtService.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	if time.Now().Unix() > data.Exp {
		return nil, errors.New("jwt token expired")
	}
	user, err := as.userService.GetUserByEmail(data.Username)
	if err != nil {
		return nil, err
	}
	return &user.UserDTO, nil
}

func (as *AuthService) AllUsers() ([]*dto.UserDTO, error) {
	userRepository, ok := as.repository.(*repository.UserRepository)
	if !ok {
		return nil, errors.New("repository is not of type *repository.UserRepository")
	}
	users, err := userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return dto.UsersArrayToDTO(users), nil
}
