package throw

func UserCreate(err error) error {
	return Error(910201, err)
}

func UserExists() error {
	return Error(910202, nil)
}

func UserError(err error) error {
	return Error(910203, err)
}

func UserNotFound() error {
	return Error(910204, nil)
}

func UserCredentialMismatch() error {
	return Error(100001, nil)
}
