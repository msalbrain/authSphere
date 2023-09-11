package server

import (
	"context"
	"database/sql"
	"embed"

	"fmt"

	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/msalbrain/authSphere/client"
	"github.com/msalbrain/authSphere/internals/api/admin"
	"github.com/msalbrain/authSphere/internals/api/auth"
	"github.com/msalbrain/authSphere/internals/database"
	service "github.com/msalbrain/authSphere/internals/service"

	// // All this to automate migration 👎
	// "github.com/golang-migrate/migrate"
	// "github.com/golang-migrate/migrate/database/sqlite3"
	// bindata "github.com/golang-migrate/migrate/source/go_bindata"
)


type EchoBuilder struct {
	e  *echo.Echo
	Q  *database.Queries
	Db *sql.DB
}

func NewEchoBuilder() *EchoBuilder {
	return &EchoBuilder{
		e: echo.New(),
	}
}

func (builder *EchoBuilder) SetupDatabase(driverName, dataSourceName , ddl string) *EchoBuilder {
	// Your database setup logic here
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}

	builder.Q = database.New(db)

	return builder
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (builder *EchoBuilder) SetupMiddlewares() *EchoBuilder {

	builder.e.Validator = &CustomValidator{validator: validator.New()}
	builder.e.Use(middleware.Recover())
	builder.e.Use(middleware.CORS())
	builder.e.Use(middleware.Gzip())
	builder.e.Use(middleware.Secure())
	builder.e.Use(middleware.Logger())

	// implement limiter
	if os.Getenv("ENABLE_LOGGER") != "" {
		builder.e.Use(middleware.Logger())
	}

	return builder

}
func (builder *EchoBuilder) SetupAPIRoutes(Env service.Config, mailins embed.FS) *EchoBuilder {

	userService := service.NewUserService(builder.Q, Env)
	mailService := service.NewSmtpMailService(mailins, builder.Q, Env)

	jwtSwitch := service.JwtOptions{
		Symmetric: service.NewSymJwt(Env.ACCESSTOKENLIFE, Env.REFRESHTOKENLIFE, Env.APPLICATION_NAME, "" , Env.JWTSecret),
		Asymmetric: service.NewAsymJwt(Env.ACCESSTOKENLIFE, Env.REFRESHTOKENLIFE, Env.APPLICATION_NAME, "", "", ""),
	}

	err := jwtSwitch.SwitchService("symmetric")

	if err != nil {
		panic(fmt.Errorf("invalid jwt option provided\nthe options available are `symmetric` or `asymmtric`"))
	}

	auth.AuthRoute(builder.e, userService, mailService, &jwtSwitch)
	admin.AdminRoute(builder.e, userService, mailService, &jwtSwitch)

	return builder
}

func (builder *EchoBuilder) SetupWebRoutes() *EchoBuilder {

	assetHandler := http.FileServer(http.FS(client.StaticFiles))
	builder.e.GET("/", echo.WrapHandler(assetHandler))
	builder.e.GET("/redoc", func(c echo.Context) error {
		return c.File("redoc-static.html")
	})

	builder.e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	return builder
}

func (builder *EchoBuilder) Build() *echo.Echo {
	return builder.e
}

func Listen(e *echo.Echo) error {
	serverHost := os.Getenv("SERVER_HOST")
	// serverPort := os.Getenv("SERVER_PORT")

	return e.Start(fmt.Sprintf("%s:%s", serverHost, "8080"))
}
