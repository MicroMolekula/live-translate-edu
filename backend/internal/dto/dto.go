package dto

type AuthDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ResponseMessageDto struct {
	Content string `json:"content"`
}

type MessageDto struct {
	User             *UserDTO `json:"user"`
	Content          string   `json:"content"`
	TranslateContent string   `json:"translate_content"`
	Room             string   `json:"room"`
}

type TranslateLanguagesDto struct {
	Source string
	Target string
}
