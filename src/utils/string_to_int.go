package utils

import (
	"log"
	"strconv"
)

func StringToInt(str string) int {
	val, err := strconv.Atoi(str)

	if err != nil {
		log.Printf("Failed to convert string to int: %v", err)
	}

	return val
}

func StringToBool(str string) bool {
	val, err := strconv.ParseBool(str)

	if err != nil {
		log.Printf("Failed to convert string to bool: %v", err)
	}

	return val
}
