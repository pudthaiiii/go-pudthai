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

	// Auth
	100001: "AUTH_CREDENTIAL_MISMATCH",
	100002: "AUTH_CREATE_JWT_TOKEN_ERROR",
	100003: "AUTH_INVALID_JWT_TOKEN",
	100004: "AUTH_REFRESH_TOKEN_ERROR",
	100005: "AUTH_ACCOUNT_TEMPORARY_LOCKED",
	100006: "AUTH_REACHED_WARNING_ATTEMPT",
	100007: "AUTH_VERIFY_AND_LOGIN_VIA_AZURE_AD_ERROR",
	100008: "AUTH_USER_NOT_FOUND_DUE_TO_AD_SETTING",

	// Roles
	910101: "ROLE_CREATE_ERROR",
	910102: "ROLE_PAGINATION_ERROR",
	910103: "ROLE_NOT_FOUND",

	// Users
	910201: "USER_CREATE_ERROR",
	910202: "USER_EMAIL_EXISTS",
	910203: "USER_ERROR",
	910204: "USER_NOT_FOUND",
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
