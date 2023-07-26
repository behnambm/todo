package utils

import (
	"log"
	"os"
	"strconv"
)

func InterfaceToString(value interface{}) string {
	// Check if the value is already a string, and return it if it is.
	if str, ok := value.(string); ok {
		return str
	}

	// Check if the value is of numeric types and convert it to a string.
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	return ""
}

func GetEnvOrPanic(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Panicln(name)
	}
	return v
}
