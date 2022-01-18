package word

import "errors"

var (
	ErrLanguageNotFound = errors.New("could not find a language")
	ErrWordNotFound     = errors.New("could not find a word")
)
