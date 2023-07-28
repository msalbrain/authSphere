package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupMiddlewares(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())

	// implement limiter
	if os.Getenv("ENABLE_LOGGER") != "" {
		e.Use(middleware.Logger())
	}
}

func apiRoutes(e *echo.Echo) {

}

func webRoutes(e *echo.Echo) {

}


func create() *echo.Echo {
	// database.SetupDatabase()

	e := echo.New()



	setupMiddlewares(e)
	apiRoutes(e)
	webRoutes(e)

	return e
}

func listen(e *echo.Echo) error {
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	return e.Start(fmt.Sprintf("%s:%s", serverHost, serverPort))
}


