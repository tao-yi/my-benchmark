package main

import (
	"log"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Get("/", func(c iris.Context) {
		c.JSON(&iris.Map{
			"message": "hello world",
		})
	})

	log.Fatal(app.Listen(":3999"))
}
