package utils

import (
	"fmt"
	"os"
)

func RequireEnv(s string, defaultValue ...string) string {
	result := os.Getenv(s)

	if result == "" {
		if len(defaultValue) == 0 || defaultValue[0] == "" {
			panic(fmt.Sprintf("error: unable to get env: %s", s))
		}

		return defaultValue[0]
	}

	return result
}
