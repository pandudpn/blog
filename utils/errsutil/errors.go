package errsutil

import "net/http"

type ResponseCode int

const (
	BadRequest ResponseCode = iota
	NotFound
	InternalServer
)

func (rc ResponseCode) Int() int {
	return [...]int{
		http.StatusBadRequest,
		http.StatusNotFound,
		http.StatusInternalServerError,
	}[rc]
}

type SystemCode int

const (
	UserNotFound SystemCode = iota
	InternalError
	PasswordNotMatch
	BodyRequired
)

func (sc SystemCode) Code() string {
	return [...]string{
		"21",
		"99",
		"22",
		"23",
	}[sc]
}

func (sc SystemCode) Message() string {
	return [...]string{
		"User not found",
		"Something went wrong",
		"Email or Password not match",
		"Invalid Parameter",
	}[sc]
}

type ErrorResponse struct {
	err        error
	message    string
	systemCode string
	statusCode int
	validation interface{}
}

// Error returns the instance of ErrorResponse
func Error(err error, sc SystemCode, rc ResponseCode, validation ...interface{}) error {
	return &ErrorResponse{
		err:        err,
		message:    sc.Message(),
		systemCode: sc.Code(),
		statusCode: rc.Int(),
		validation: validation,
	}
}

// Error message from ErrorResponse instance
func (er *ErrorResponse) Error() string {
	return er.err.Error()
}

// Message return reason an error from ErrorResponse instance
func (er *ErrorResponse) Message() string {
	return er.message
}

// SystemCode return a system code for track the error from ErrorResponse instance
func (er *ErrorResponse) SystemCode() string {
	return er.systemCode
}

// StatusCode return Status Code Response from ErrorResponse instance
func (er *ErrorResponse) StatusCode() int {
	return er.statusCode
}

// Validation return a field validation error from ErrorResponse instance
func (er *ErrorResponse) Validation() interface{} {
	return er.validation
}
