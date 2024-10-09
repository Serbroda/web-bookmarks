package random

import (
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	CharsetAlphaLower        = "abcdefghijklmnopqrstuvwxyz"
	CharsetAlphaUpper        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetNumbers           = "0123456789"
	CharsetSpecialCharacters = "-_.$@"
	CharsetAlphaNumeric      = CharsetAlphaLower + CharsetAlphaUpper + CharsetNumbers
)

func RandomNumber(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func RandomString(length int) string {
	return RandomStringWithCharset(length, CharsetAlphaNumeric)
}

func RandomStringWithCharset(length int, letters string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[seededRand.Intn(len(letters))]
	}
	return string(b)
}
