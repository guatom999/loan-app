package handlers

import "github.com/labstack/echo/v4"

type (
	HandlerInterface interface {
		Loan(c echo.Context) error
	}
)
