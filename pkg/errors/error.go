package errors

import (
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("not found")

func HttpErrCode(err error) int {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
