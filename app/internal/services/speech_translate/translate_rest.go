package speech_translate

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type TranslateRequest struct {
	TargetLanguage string `json:"targetLanguageCode"`
	SourceLanguage string `json:"sourceLanguageCode"`
	Text           string `json:"texts"`
}

type TextTranslate struct {
	Text string `json:"text"`
}

type TranslateResponse struct {
	Translations []TextTranslate `json:"translations"`
}

func TranslateRest(text string) (string, error) {
	reqBody := TranslateRequest{
		TargetLanguage: "en",
		SourceLanguage: "ru",
		Text:           text,
	}
	url := "https://translate.api.cloud.yandex.net/translate/v2/translate"

	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// Создаем новый запрос
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}

	// Устанавливаем заголовок с типом данных в теле запроса
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Api-Key ")

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseJson, err := io.ReadAll(resp.Body)
	var res = &TranslateResponse{}
	err = json.Unmarshal(responseJson, res)
	if err != nil {
		return "", err
	}
	return res.Translations[0].Text, nil
}
