package main

import (
	"api/internal"
	"api/users"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var server = flag.String("server", ":1323", "Host") 
var url = flag.String("url", "mongodb://user:pass@167.172.76.207:27017", "Host") 

func main() {
	fmt.Println(os.Getenv("SERVER"))
	flag.Parse()
	// Echo server
	e := echo.New()
	// Add middlewares

	// Routers
	e.GET("/", homeHandler)

	// Users
	client := internal.NewMongoClient(*url)
	repo := users.NewUserRepository(client)
	service := users.NewUserService(&repo)
	e.GET("/users", users.GetUserHandler(service))

	// Start server
	e.Logger.Fatal(e.Start(*server))
}

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
