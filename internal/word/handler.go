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
	g := h.f.Group("/v1/word")

	g.Get("/random", h.GetRandomWord)
	g.Post("/", h.CreateWord)
	g.Post("/:id/:code/translate", h.CreateTranslate)
}

func (h *handler) GetRandomWord(f *fiber.Ctx) error {
	word, err := h.s.GetRandomWord()

	if err != nil {
		return f.
			Status(fiber.StatusConflict).
			SendString(err.Error())
	}

	return f.JSON(word)
}

func (h *handler) CreateWord(f *fiber.Ctx) error {
	word := new(CreateNewWordDTO)

	if err := f.BodyParser(&word); err != nil {
		return f.
			Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	err := h.s.CreateNewWord(word)

	if err != nil {
		return f.
			Status(fiber.StatusConflict).
			SendString(err.Error())
	}

	return f.SendString("OK")
}

func (h *handler) CreateTranslate(f *fiber.Ctx) error {
	translate := new(CreateNewTranslateDTO)
	wordID := f.Params("id")
	languageCode := f.Params("code")

	if err := f.BodyParser(&translate); err != nil {
		return f.
			Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	err := h.s.CreateNewTranslate(wordID, languageCode, translate)

	if err != nil {
		return f.
			Status(fiber.StatusConflict).
			SendString(err.Error())
	}

	return f.SendString("OK")
}
