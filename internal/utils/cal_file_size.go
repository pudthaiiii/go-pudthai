package utils

import (
	"errors"
	"strconv"
	"strings"
)

func CalFileSize(sizeStr string) (int64, error) {
	const (
		mb = 1024 * 1024
		kb = 1024
	)

	sizeStr = strings.ToLower(strings.TrimSpace(sizeStr))
	var size int64
	var err error

	if strings.HasSuffix(sizeStr, "mb") {
		sizeStr = strings.TrimSuffix(sizeStr, "mb")
		size, err = strconv.ParseInt(sizeStr, 10, 64)
		if err != nil {
			return 0, err
		}
		size *= mb
	} else if strings.HasSuffix(sizeStr, "kb") {
		sizeStr = strings.TrimSuffix(sizeStr, "kb")
		size, err = strconv.ParseInt(sizeStr, 10, 64)
		if err != nil {
			return 0, err
		}
		size *= kb
	} else if strings.HasSuffix(sizeStr, "b") {
		sizeStr = strings.TrimSuffix(sizeStr, "b")
		size, err = strconv.ParseInt(sizeStr, 10, 64)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("invalid size format. Use 'b', 'kb', or 'mb'")
	}

	return size, nil
}
