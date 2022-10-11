package main

import (
	"api/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo server
	e := echo.New()
	// Add middlewares

	// Routers
	e.GET("/", homeHandler)
	e.GET("/users", users.UserHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}


