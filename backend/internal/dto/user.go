package dto

import "github.com/live-translate-edu/internal/models"

type UserDTO struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Role    string `json:"role"`
}

type UserCreateDTO struct {
	UserDTO
	Password string `json:"password"`
}

type UserFullDTO struct {
	UserDTO
	Password string `json:"password"`
}

func UserFullDTOToModel(dto *UserFullDTO) *models.User {
	return &models.User{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Email:    dto.Email,
		Role:     dto.Role,
		Password: dto.Password,
	}
}

func UserModelToFullDTO(model *models.User) *UserFullDTO {
	return &UserFullDTO{
		UserDTO:  *UserToDTO(model),
		Password: model.Password,
	}
}

func UserCreateDTOToModel(dto *UserCreateDTO) *models.User {
	return &models.User{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Email:    dto.Email,
		Role:     dto.Role,
		Password: dto.Password,
	}
}

func UserToDTO(model *models.User) *UserDTO {
	return &UserDTO{
		Id:      model.ID,
		Name:    model.Name,
		Surname: model.Surname,
		Email:   model.Email,
		Role:    model.Role,
	}
}

func UsersArrayToDTO(models []*models.User) []*UserDTO {
	arrayDto := make([]*UserDTO, len(models))
	for i := 0; i < len(models); i++ {
		arrayDto[i] = UserToDTO(models[i])
	}
	return arrayDto
}

func UserDTOToModel(user *UserDTO) *models.User {
	return &models.User{
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Role:    user.Role,
	}
}
