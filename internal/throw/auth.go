package throw

func UserCredentialMismatch() error {
	return Error(100001, nil)
}

func GenerateJwtTokenError(err error) error {
	return Error(100002, err)
}
