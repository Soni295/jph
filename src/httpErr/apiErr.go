package httperr

import (
	"errors"
	"net/http"
)

type APIError struct {
	Status  int
	Message string
}

func FormError(err error) APIError {
	servErr := Error{}

	if !errors.As(err, &servErr) {
		return APIError{Message: "unknow error",
			Status: http.StatusInternalServerError,
		}
	}

	message := servErr.AppErr().Error()
	status := 500

	switch servErr.ServErr() {
	case ErrBadRequest:
		status = http.StatusBadRequest
	}

	return APIError{
		Message: message,
		Status:  status,
	}
}