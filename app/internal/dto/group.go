package dto

import "github.com/live-translate-edu/internal/models"

type Group struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Code  string `json:"code"`
}

func GroupToModel(group *Group) *models.Group {
	return &models.Group{
		Title: group.Title,
		Code:  group.Code,
	}
}
