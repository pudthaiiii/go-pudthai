package utils

import (
	"fmt"
	"regexp"
)

func FilterThrowExceptions(message string) []string {
	re := withErrors()

	matches := re.FindStringSubmatch(message)
	if matches == nil {
		re = withOutErrors()
		matches = re.FindStringSubmatch(message)
	}

	return matches
}

func withErrors() *regexp.Regexp {
	re, err := regexp.Compile(`\[(\d+)\]: ([^:]+): (.+)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	return re
}

func withOutErrors() *regexp.Regexp {
	re, err := regexp.Compile(`\[(\d+)\]: (.+):`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	return re
}
