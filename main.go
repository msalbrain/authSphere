package main

import (
	"log"
	_ "embed"
	"embed"
	"fmt"
	"github.com/msalbrain/authSphere/internals/server"
	service "github.com/msalbrain/authSphere/internals/service"
)

//go:embed migrations/0001_init_db.up.sql
var ddl string


//go:embed mail/*
var embeddedTemplates embed.FS

func main() {

	config := service.GetEnvirometConfig()

	dbString := fmt.Sprintf("%s?mode=%s", config.Dbname, config.Mode)

	app := server.NewEchoBuilder().
		SetupDatabase("sqlite3", dbString, ddl).
		SetupMiddlewares().
		SetupAPIRoutes(config, embeddedTemplates).
		SetupWebRoutes().
		Build()


	err := server.Listen(app)

	log.Fatal(err)
}
