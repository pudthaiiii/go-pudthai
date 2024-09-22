package throw

func UserExists() error {
	return Error(910202, nil)
}

func UserError(err error) error {
	return Error(910203, err)
}

func UserCreate(err error) error {
	return Error(910201, err)
}
