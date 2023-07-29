package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminRoute(e *echo.Echo) {

	g := e.Group("/admin")

	g.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Param("id"))
	})

}
