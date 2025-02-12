package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/live-translate-edu/internal/dto"
	"time"
)

var (
	ErrorInvalidJWTToken = errors.New("invalid JWT token")
)

type DataJwt struct {
	Id       uint   `json:"id"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTService struct {
	jwtSecret string
}

func NewJwtService(jwtSecret string) *JWTService {
	return &JWTService{jwtSecret: jwtSecret}
}

func (s *JWTService) GenerateTokenByUser(user *dto.UserFullDTO, ttl int) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, DataJwt{
		Exp:      time.Now().Add(time.Duration(ttl) * time.Minute).Unix(),
		Iat:      time.Now().Unix(),
		Username: user.Email,
		Role:     user.Role,
		Id:       user.Id,
	})

	token, err := jwtToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *JWTService) ParseToken(tokenString string) (*DataJwt, error) {
	token, err := jwt.ParseWithClaims(tokenString, &DataJwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	result, ok := token.Claims.(*DataJwt)
	if !ok {
		return nil, ErrorInvalidJWTToken
	}

	return result, nil
}
