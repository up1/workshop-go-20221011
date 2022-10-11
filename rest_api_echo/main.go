package main

import (
	"api/users"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var server = flag.String("server", ":1323", "Host")

func main() {
	fmt.Println(os.Getenv("SERVER"))
	flag.Parse()
	// Echo server
	e := echo.New()
	// Add middlewares

	// Routers
	e.GET("/", homeHandler)
	e.GET("/users", users.UserHandler)

	// Start server
	e.Logger.Fatal(e.Start(*server))
}

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

