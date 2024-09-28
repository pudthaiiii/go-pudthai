package throw

func MerchantNotFound() error {
	return Error(910301, nil)
}
