package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"

	service "github.com/msalbrain/authSphere/internals/service"
)

func AdminRoute(e *echo.Echo, u service.UserService, mail service.MailService, jwtoption service.JwtOptionService) {

	g := e.Group("/admin")

	g.GET("/users/:id", GreetAdmin)

}



/*
@title Swagger Example API
@version 1.0
@description This is a sample server Petstore server.
@termsOfService http://swagger.io/terms/
*/
func GreetAdmin(c echo.Context) error {

	return c.String(http.StatusOK, c.Param("id"))
}
