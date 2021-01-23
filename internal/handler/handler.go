package handler

import "github.com/labstack/echo/v4"

// New wraps every routes
func New(e *echo.Echo) {
	public := e.Group("/public")

	user := public.Group("/user")
	user.GET("/test", Test)
}
