package database

import (
	"fmt"
	"gora/config"
	"gora/internal/word"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseClient(conf *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			conf.Database.Hostname,
			conf.Database.Username,
			conf.Database.Password,
			conf.Database.Name,
			conf.Database.Port,
			conf.Database.SSL,
		),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&word.Language{}, &word.Word{}, &word.Dictionary{})

	// Seed
	CreateLanguages(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}
