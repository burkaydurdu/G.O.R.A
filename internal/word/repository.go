package word

import "gorm.io/gorm"

type Repository interface {
	GetRandomWord() (*Word, error)
	CreateNewWord(word *Word) error
	CreateNewTranslate(dictionary *Dictionary) error
	GetLanguageByCode(code string) (*Language, error)
	GetWordByID(id string) (*Word, error)
	GetRandomTranslate(fromID, toID string) (*Dictionary, error)
}

type repository struct {
	g *gorm.DB
}

func NewRepository(g *gorm.DB) Repository {
	r := &repository{
		g: g,
	}

	return r
}

func (r *repository) GetRandomWord() (word *Word, err error) {
	result := r.g.Joins("Language").First(&word)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return word, nil
}

func (r *repository) CreateNewWord(word *Word) (err error) {
	result := r.g.Create(&word)

	return result.Error
}

func (r *repository) GetLanguageByCode(code string) (language *Language, err error) {
	result := r.g.First(&language, "code = ?", code)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrLanguageNotFound
	}

	return language, nil
}

func (r *repository) GetWordByID(id string) (word *Word, err error) {
	result := r.g.First(&word, id)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrWordNotFound
	}

	return word, nil
}

func (r *repository) CreateNewTranslate(dictionary *Dictionary) error {
	result := r.g.Create(&dictionary)

	return result.Error
}

func (r *repository) GetRandomTranslate(fromID, toID string) (dictionary *Dictionary, err error) {
	result := r.g.
		Joins("Language").
		Joins("Word").
		Where("\"Word\".language_id = ?", fromID).
		Where("\"dictionaries\".language_id = ?", toID).
		First(&dictionary)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return dictionary, nil
}
