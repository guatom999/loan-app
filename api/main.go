package main

import (
	"net/http"

	"github.com/finance-app/config"
	"github.com/finance-app/databases"
	"github.com/finance-app/pkg/handlers"
	"github.com/finance-app/pkg/repositories"
	"github.com/finance-app/pkg/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	LoadRequest struct {
		FullName string `json:"fullname"`
		Salary   int64  `json:"salary"`
	}
)

func main() {

	cfg := config.GetConfig()

	db := databases.ConnDb(cfg)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	repo := repositories.NewRepository(db)
	usecase := usecases.NewUseCase(repo)
	handlers := handlers.NewHandler(usecase)

	route := e.Group("/api/vi")

	route.POST("/loan", handlers.Loan)

	e.Logger.Fatal(e.Start(cfg.App.Port))
}
