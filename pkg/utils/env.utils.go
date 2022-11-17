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

func GetEnv(key string) (string, bool) {
	once.Do(func() {
		godotenv.Load()
	})
	return os.LookupEnv(key)
}

func GetEnvFallback(key, fallback string) string {
	if value, ok := GetEnv(key); ok {
		return value
	}
	return fallback
}

func MustGetEnv(key string) string {
	if value, ok := GetEnv(key); ok {
		return value
	}
	panic("Mandatory env " + key + " not found")
}

func MustParseInt64(value string) int64 {
	val, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		panic("Failed to parse " + value + " to int64")
	}
	return int64(val)
}
