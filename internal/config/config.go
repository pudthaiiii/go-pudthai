package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	cfg map[string]interface{}
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		cfg: make(map[string]interface{}),
	}, nil
}

func (c *Config) Get(key string) map[string]interface{} {
	if _, exists := c.cfg[key]; !exists {
		return nil
	}

	return c.cfg[key].(map[string]interface{})
}

func (c *Config) Add(key string, value map[string]interface{}) {
	c.cfg[key] = value
}

func (c *Config) Initialize() {
	c.loadAppConfig()
}

func getEnv(key string, defaultValue ...string) string {
	var defaultVal string

	if len(defaultValue) > 0 {
		defaultVal = defaultValue[0]
	}

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
