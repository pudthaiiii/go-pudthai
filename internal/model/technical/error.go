package technical

import "errors"

type ErrorCause struct {
	Cause      string `json:"cause"`
	Err        error  `json:"error"`
	StatusCode int    `json:"statusCode"`
}

// Error type
const (
	errorIncorrectParameter   = "ERROR_INCORRECT_PARAMETER"
	errorIncorrectPayload     = "ERROR_INCORRECT_PAYLOAD"
	errorNoProductPermission  = "ERROR_NO_PRODUCT_PERMISSION"
	errorMethodNotAllow       = "ERROR_METHOD_NOT_ALLOW"
	errorDatabaseTypeMismatch = "ERROR_DATABASE_TYPE_MISMATCH"
	errorExceedQuantity       = "ERROR_EXCEED_QUANTITY"
	errorDataNotFound         = "ERROR_DATA_NOT_FOUND"
	errorOtherDocChanged      = "ERROR_OTHER_DOC_CHANGED"
	errorItemCancelled        = "ERROR_ITEM_HAS_BEEN_CANCELLED"
	errorItemNoRowAffected    = "ERROR_ITEM_NO_ROW_AFFECTED"
	errorItemApproved         = "ERROR_ITEM_HAS_BEEN_APPROVED"
	errorNoPermission         = "ERROR_NO_PERMISSION"
	errorWaitingOther         = "ERROR_WAITING_OTHER"
	errorAlreadyExistData     = "ERROR_ALREADY_EXIST_DATA"
)

type staticErrorCause struct{}

var (
	sec *staticErrorCause
)

func StaticErrorCause() *staticErrorCause {
	return sec
}

func (sec *staticErrorCause) GetErrorIncorrectPayload() error {
	return errors.New(errorIncorrectPayload)
}

func (sec *staticErrorCause) GetErrorMethodNotAllow() error {
	return errors.New(errorMethodNotAllow)
}

func (sec *staticErrorCause) GetErrorNoPermission() error {
	return errors.New(errorNoPermission)
}

func (sec *staticErrorCause) GetErrorDataNotFound() error {
	return errors.New(errorDataNotFound)
}

func (sec *staticErrorCause) GetErrorCause(cause string, err error, code int) ErrorCause {
	return ErrorCause{
		Cause:      cause,
		Err:        err,
		StatusCode: code,
	}
}
