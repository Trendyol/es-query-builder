package errorx

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	Details string
	Status  int
}

func (e Error) Error() string {
	return fmt.Sprintf("status code: %d, details: %s", e.Status, e.Details)
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Status == http.StatusNotFound
	}
	return false
}
