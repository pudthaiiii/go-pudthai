package throw

import (
	"fmt"
	"go-ibooking/src/pkg/logger"
)

var ErrorCodes = map[int]string{
	0:      "UNDEFINED_ERROR",
	900400: "BAD_REQUEST",
	900401: "UNAUTHORIZED",
	900403: "FORBIDDEN",
	900422: "VALIDATE_ERROR",
	910001: "TYPEORM_ERROR",
	910002: "INSUFFICIENT_ABILITY",
	910003: "UPLOADER_ERROR",

	// Roles
	910101: "ROLE_ERROR",
	910102: "ROLE_NOT_FOUND",
}

func Error(code int, err error) error {
	if err == nil {
		msg := fmt.Sprintf("[%d]: %s:", code, ErrorCodes[code])
		logger.Log.Error().Msg(msg)
		return fmt.Errorf(msg)
	}

	msg := fmt.Sprintf("[%d]: %s: %s", code, ErrorCodes[code], err.Error())
	logger.Log.Err(err).Msg(msg)
	return fmt.Errorf(msg)
}
