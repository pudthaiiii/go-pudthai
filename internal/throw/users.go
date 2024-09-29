package throw

func UserCreate(err error) error {
	return Error(910201, err, 500)
}

func UserExists() error {
	return Error(910202, nil)
}

func UserError(err error) error {
	return Error(910203, err, 500)
}

func UserNotFound() error {
	return Error(910204, nil, 404)
}
