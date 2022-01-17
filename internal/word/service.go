package word

import "gorm.io/gorm"

type Service interface {
	GetRandomWord() (*Word, error)
	CreateNewWord(newWord *CreateNewWordDTO) error
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

	word.Language = Language{Name: newWord.Language, Model: gorm.Model{ID: 2}}
	word.Content = newWord.Word

	err := s.r.CreateNewWord(word)

	return err
}
