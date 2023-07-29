package main

import (
	"github.com/msalbrain/authSphere/internals/server"
	"log"
)


func main() {
	app := server.Create()

	err := server.Listen(app)

	log.Fatal(err)
}

