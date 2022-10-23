package utils

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
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
