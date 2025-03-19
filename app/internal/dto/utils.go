package dto

import "strconv"

type DataForFormLesson struct {
	Languages []Option `json:"languages"`
	Groups    []Option `json:"groups"`
}

type Option struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func CreateDataForFormLesson(languages []Language, groups []Group) DataForFormLesson {
	languageOptions := make([]Option, len(languages))
	for i, language := range languages {
		languageOptions[i] = Option{
			Id:    language.Code,
			Title: language.Title,
		}
	}
	groupOptions := make([]Option, len(groups))
	for i, group := range groups {
		groupOptions[i] = Option{
			Id:    strconv.Itoa(group.Id),
			Title: group.Code,
		}
	}
	return DataForFormLesson{
		Languages: languageOptions,
		Groups:    groupOptions,
	}
}
