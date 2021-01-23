package util

import (
	"net/http"
)

var errorMap = map[error]int{}

// MapError determine which code should be returned
func MapError(err error) int {
	for i := range errorMap {
		if i == err {
			return errorMap[i]
		}
	}
	return http.StatusInternalServerError
}
