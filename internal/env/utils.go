package env

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnvVar(key string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}

	value := os.Getenv(key)
	if len(value) == 0 {
		return "", nil
	}

	return value, nil
}
