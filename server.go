package main

import (
	"fmt"
	"gora/config"
	"gora/internal/middleware"
	"gora/internal/persistent/database"
	"gora/internal/word"

	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	f *fiber.App
	c *config.Config
}

func NewServer(config *config.Config) (server *Server) {
	app := fiber.New()

	// Set Middleware
	app.Use(middleware.NewAuthMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/dashboard", monitor.New())

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
