package word

import "gorm.io/gorm"

type Repository interface {
	GetRandomWord() (*Word, error)
	CreateNewWord(word *Word) error
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
	result := r.g.First(&word)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return word, nil
}

func (r *repository) CreateNewWord(word *Word) (err error) {
	result := r.g.Create(&word)

	return result.Error
}
