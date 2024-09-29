package utils

import (
	"encoding/json"
	"fmt"
)

func MarshalIndent(v any) string {
	stringData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal role:", err)
	}

	return string(stringData)
}
