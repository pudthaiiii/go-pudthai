package utils

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func Ms(durationStr string) (int64, error) {
	return parseDuration(durationStr)
}

func parseDuration(durationStr string) (int64, error) {
	re := regexp.MustCompile(`^(\d+)([smhd])$`)
	matches := re.FindStringSubmatch(durationStr)
	if len(matches) != 3 {
		return 0, errors.New("invalid duration format")
	}

	value, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	unit := matches[2]
	var duration time.Duration

	switch unit {
	case "s":
		duration = time.Duration(value) * time.Second
	case "m":
		duration = time.Duration(value) * time.Minute
	case "h":
		duration = time.Duration(value) * time.Hour
	case "d":
		duration = time.Duration(value) * 24 * time.Hour
	default:
		return 0, errors.New("unknown time unit")
	}

	return int64(duration.Seconds()), nil
}
