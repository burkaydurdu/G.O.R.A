package database

import (
	"gora/internal/word"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

func CreateLanguages(db *gorm.DB) {
	languageEN := word.Language{
		Name: "English",
		Code: "EN",
	}

	languageTR := word.Language{
		Name: "Türkçe",
		Code: "TR",
	}

	languages := []word.Language{languageEN, languageTR}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&languages)
}
