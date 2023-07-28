package main

import "fmt"

// "net/http"

// "github.com/labstack/echo/v4"


type dog struct {
	name string
	alive bool
}

func (d dog) getName() string {
	return d.name
}

func (d dog) getDogState() bool {
	return d.alive
}

type cat struct {
	name string
	color string
}

// func (c cat) getName() string {
// 	return c.name
// }

func (c cat) getCatColor() string {
	return c.color
}


type Animal interface {
	getName() string
	getDogState() bool
}


func letSee(animal Animal) {

	fmt.Println(animal)
}

func main() {
	d := dog{
		name: "bingo",
		alive: false,
	}

	c := cat{
		name: "garfield",
		color: "orange",
	}


	letSee(d)
	letSee(c)

	// e := echo.New()

    // e.GET("/", func(c echo.Context) error {
    //     return c.String(http.StatusOK, "Hello, World!")
    // })

    // e.Logger.Fatal(e.Start(":1323"))
}
