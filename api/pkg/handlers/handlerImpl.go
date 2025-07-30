package handlers

import (
	"net/http"

	"github.com/finance-app/pkg"
	"github.com/finance-app/pkg/usecases"
	"github.com/labstack/echo/v4"
)

type (
	handler struct {
		usecase usecases.UsecaseInterface
	}
)

func NewHandler(usecase usecases.UsecaseInterface) HandlerInterface {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) Loan(c echo.Context) error {

	req := new(pkg.UserCreateStateMentRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "statement invalid")
	}

	if err := h.usecase.CheckStateMent(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "approve")
}
