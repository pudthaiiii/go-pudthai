package throw

func UploadError(err error) error {
	return Error(910003, err)
}
