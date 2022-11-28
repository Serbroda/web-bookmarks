package utils

import (
	"fmt"
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
		fmt.Printf("%s=%v\n", key, value)
		return value
	}
	fmt.Printf("%s (fallback)=%v\n", key, fallback)
	return fallback
}

func MustGetEnv(key string) string {
	if value, ok := GetEnv(key); ok {
		fmt.Printf("%s=%v\n", key, value)
		return value
	}
	panic("Mandatory env " + key + " not found")
}

func MustParseInt64(value string) int64 {
	val, err := ParseInt64(value)
	if err != nil {
		panic("Failed to parse " + value + " to int64")
	}
	return val
}

func ParseInt64(value string) (int64, error) {
	val, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return -1, err
	}
	return int64(val), nil
}
