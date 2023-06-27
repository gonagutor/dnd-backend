package utils

import (
	"math/rand"
)

const (
	POSSIBLE_CHARACTERS = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenerateRandomCode(count int) string {
	b := make([]byte, count)
	for i := range b {
		b[i] = POSSIBLE_CHARACTERS[rand.Intn(len(POSSIBLE_CHARACTERS))]
	}
	return string(b)
}
