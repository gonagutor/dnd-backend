package validators

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

const (
	MIN_LENGTH = 8
	MIN_NUMBER = 1
	MIN_UPPER  = 1
	MIN_LOWER  = 1
	MIN_SYMBOL = 1
)

func Password(fl validator.FieldLevel) bool {
	letters := 0
	number := 0
	upper := 0
	lower := 0
	symbol := 0

	for _, c := range fl.Field().String() {
		letters++
		switch {
		case unicode.IsNumber(c):
			number++
		case unicode.IsLower(c):
			lower++
		case unicode.IsUpper(c):
			upper++
		case unicode.IsPunct(c) || unicode.IsSymbol(c) || c == ' ':
			symbol++
		}
	}

	minLength := letters >= MIN_LENGTH
	minNumber := number >= MIN_NUMBER
	minUpper := upper >= MIN_UPPER
	minLower := lower >= MIN_LOWER
	minSymbol := symbol >= MIN_SYMBOL

	return (minLength && minNumber && minUpper && minLower && minSymbol)
}
