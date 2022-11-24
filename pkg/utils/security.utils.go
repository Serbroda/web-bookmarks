package utils

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const CharsetAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Charset = CharsetAlpha + "0123456789_-"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandomString(n int) string {
	return RandomStringWithCharset(n, Charset)
}

func RandomStringWithCharset(n int, letters string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[seededRand.Intn(len(letters))]
	}
	return string(b)
}
