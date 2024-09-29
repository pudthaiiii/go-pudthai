package throw

func UserCredentialMismatch() error {
	return Error(100001, nil, 401)
}

func GenerateJwtTokenError(err error) error {
	return Error(100002, err, 500)
}

func InvalidJwtToken(err error) error {
	return Error(100003, err, 401)
}
