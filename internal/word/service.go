package word

import "strings"

type Service interface {
	GetRandomWord() (*Word, error)
	CreateNewWord(newWord *CreateNewWordDTO) error
	CreateNewTranslate(wordID string, languageCode string, newTranslate *CreateNewTranslateDTO) error
	GetRandomTranslate(fromID, toID string) (*Dictionary, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) GetRandomWord() (*Word, error) {
	word, err := s.r.GetRandomWord()

	return word, err
}

func (s *service) CreateNewWord(newWord *CreateNewWordDTO) error {
	word := new(Word)

	language, err := s.r.GetLanguageByCode(newWord.LanguageCode)

	if err == ErrLanguageNotFound {
		return err
	}

	word.Language = *language
	word.Content = newWord.Word

	err = s.r.CreateNewWord(word)

	return err
}

func (s *service) CreateNewTranslate(wordID string, languageCode string, newTranslate *CreateNewTranslateDTO) error {
	word, werr := s.r.GetWordByID(wordID)

	if werr == ErrWordNotFound {
		return werr
	}

	language, lerr := s.r.GetLanguageByCode(strings.ToUpper(languageCode))

	if lerr == ErrLanguageNotFound {
		return lerr
	}

	dictionary := new(Dictionary)
	dictionary.Word = *word
	dictionary.Language = *language
	dictionary.Content = newTranslate.Content

	err := s.r.CreateNewTranslate(dictionary)

	return err
}

func (s *service) GetRandomTranslate(fromID, toID string) (*Dictionary, error) {
	translate, err := s.r.GetRandomTranslate(fromID, toID)

	return translate, err
}
