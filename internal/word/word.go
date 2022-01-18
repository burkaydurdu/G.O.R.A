package word

import (
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_lg"`
	Code string `gorm:"uniqueIndex:idx_cd"`
}

type Word struct {
	gorm.Model
	Language   Language
	LanguageID uint         `gorm:"uniqueIndex:idx_wd"`
	Content    string       `gorm:"uniqueIndex:idx_wd"`
	Dictionary []Dictionary `gorm:"foreignKey:WordID"`
}

type Dictionary struct {
	gorm.Model
	WordID     uint `gorm:"uniqueIndex:idx_dc"`
	Word       Word
	Language   Language
	LanguageID uint `gorm:"uniqueIndex:idx_dc"`
	Content    string
}
