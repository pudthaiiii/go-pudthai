package throw

func ValidateRecaptchaError() error {
	return Error(910004, nil)
}

func RecaptchaError() error {
	return Error(910005, nil, 500)
}
