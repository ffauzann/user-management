package util

import (
	"net/http"

	"example.com/user-management/internal/response"
	"github.com/labstack/echo/v4"
)

// SuccessResponse handle every 200 http code response
func SuccessResponse(c echo.Context, msg string, data interface{}) error {
	i := response.SuccessResponse{
		Base: response.Base{
			Code:    http.StatusOK,
			Message: msg,
		},
		Data: data,
	}

	return c.JSON(http.StatusOK, i)
}

// ErrorResponse determine which error should be thrown as a response
func ErrorResponse(c echo.Context, err error) error {
	code := MapError(err)
	switch code {
	case http.StatusBadRequest:
		return HandleBadRequest(c, err.Error())
	}
	return HandleInternalServerError(c, err.Error())
}

// HandleBadRequest handle every 400 http code response
func HandleBadRequest(c echo.Context, msg string) error {
	data := response.Base{
		Code:    http.StatusBadRequest,
		Message: msg,
	}

	return c.JSON(http.StatusBadRequest, data)
}

// HandleInternalServerError handle every 500 http code response
func HandleInternalServerError(c echo.Context, msg string) error {
	data := response.Base{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}

	return c.JSON(http.StatusInternalServerError, data)
}
