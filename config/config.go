package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	BodyLimit    int
	ReadTimeout  int
	WriteTimeout int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	bodyLimit, err := strconv.Atoi(os.Getenv("BODY_LIMIT"))
	if err != nil {
		bodyLimit = 100
	}

	readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		readTimeout = 10
	}

	writeTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		writeTimeout = 10
	}

	return &Config{
		Port:         getEnv("PORT", "3000"),
		BodyLimit:    bodyLimit,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
