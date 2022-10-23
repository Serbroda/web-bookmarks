package utils

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once
)

func GetEnv(key, fallback string) string {
	once.Do(func() {
		godotenv.Load()
	})

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func MustParseInt64(value string) int64 {
	val, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		panic("Failed to parse")
	}
	return int64(val)
}
