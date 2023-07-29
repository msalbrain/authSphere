package server

import (
	"fmt"
	"os"
	"net/http"


	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/msalbrain/authSphere/internals/api/admin"
	"github.com/msalbrain/authSphere/client"
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
	admin.AdminRoute(e)
}

func webRoutes(e *echo.Echo) {

	assetHandler := http.FileServer(http.FS(client.StaticFiles))
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	// e.File("/web", "C:/Users/salman/Desktop/go-go/authSphere/client/dist/index.html")
}


func Create() *echo.Echo {
	// database.SetupDatabase()

	e := echo.New()

	setupMiddlewares(e)
	apiRoutes(e)
	webRoutes(e)

	return e
}

func Listen(e *echo.Echo) error {
	serverHost := os.Getenv("SERVER_HOST")
	// serverPort := os.Getenv("SERVER_PORT")

	return e.Start(fmt.Sprintf("%s:%s", serverHost, "8080"))
}


