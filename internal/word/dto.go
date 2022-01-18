package word

type CreateNewWordDTO struct {
	Word         string `json:"word"`
	LanguageCode string `json:"language_code"`
}

type CreateNewTranslateDTO struct {
	Content string `json:"content"`
}
