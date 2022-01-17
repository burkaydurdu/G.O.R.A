package word

import (
	"github.com/gofiber/fiber/v2"
)

type HTTPHandler interface {
	GetRandomWord(ctx *fiber.Ctx) error
}

type handler struct {
	f *fiber.App
	s Service
}

func NewHandler(f *fiber.App, s Service) HTTPHandler {
	h := &handler{
		f: f,
		s: s,
	}

	h.registerRoutes()

	return h
}

func (h *handler) registerRoutes() {
	g := h.f.Group("/random")

	g.Get("/en/word", h.GetRandomWord)
	g.Post("/en/word", h.CreateWord)
}

func (h *handler) GetRandomWord(f *fiber.Ctx) error {
	word, err := h.s.GetRandomWord()

	if err != nil {
		return f.
			Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return f.JSON(word)
}

func (h *handler) CreateWord(f *fiber.Ctx) error {
	word := new(CreateNewWordDTO)

	if err := f.BodyParser(&word); err != nil {
		return err
	}

	err := h.s.CreateNewWord(word)

	if err != nil {
		return err
	}

	return f.SendString("OK")
}
