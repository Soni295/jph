package httperr

import "errors"

var (
	ErrBadRequest          error = errors.New("Bad Request")
	ErrInternalFailure           = errors.New("Internal Failure")     // 500
	ErrNotFound                  = errors.New("Not Found")            // 404
	ErrUnprocessableEntity       = errors.New("Unprocessable Entity") // 400
)

func NewError(servErr, appErr error) error {
	return &Error{
		servErr: servErr,
		appErr:  appErr,
	}
}

type Error struct {
	appErr  error
	servErr error
}

func (e Error) Error() string {
	return errors.Join(e.servErr, e.appErr).Error()
}

func (e Error) AppErr() error  { return e.appErr }
func (e Error) ServErr() error { return e.servErr }
func (e Error) Status() int {
	switch e.appErr {
	case ErrBadRequest:
		return 400
	case ErrInternalFailure:
		return 500
	case ErrNotFound:
		return 404
	case ErrUnprocessableEntity:
		return 400
	}

	return 500
}