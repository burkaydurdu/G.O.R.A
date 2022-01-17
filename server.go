package main

import (
	"fmt"
	"gora/config"
	"gora/internal/persistent/database"
	"gora/internal/word"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	f *fiber.App
	c *config.Config
}

func NewServer(config *config.Config) (server *Server) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	server = &Server{
		f: app,
		c: config,
	}

	return server
}

func (s *Server) Start() error {
	db, err := database.NewDatabaseClient(s.c)

	if err != nil {
		panic(err)
	}

	wordRepository := word.NewRepository(db)
	wordService := word.NewService(wordRepository)
	word.NewHandler(s.f, wordService)

	return s.f.Listen(fmt.Sprintf(":%d", s.c.Server.Port))
}
