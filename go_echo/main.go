package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":3999"))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, &echo.Map{
		"message": "hello world",
	})
}
