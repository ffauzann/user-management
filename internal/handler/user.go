package handler

import (
	"example.com/user-management/internal/util"
	"github.com/labstack/echo/v4"
)

// Test handle testing route
func Test(c echo.Context) (err error) {
	return util.SuccessResponse(c, "testing route", nil)
}
