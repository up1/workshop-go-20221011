package main

import (
	"api/internal"
	"api/users"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

var server = flag.String("server", ":1323", "Host") 
var url = flag.String("url", "mongodb://user:pass@167.172.76.207:27017", "Host") 

func main() {

	// Init tracer
	tp, err := internal.InitTracer("http://zipkin:9411/api/v2/spans", "demo-api")
	if err != nil {
		log.Fatal(err)
	}
	defer func ()  {
		tp.Shutdown(context.Background())
	}()

	fmt.Println(os.Getenv("SERVER"))
	flag.Parse()
	// Echo server
	e := echo.New()
	// Add global middlewares
	p := prometheus.NewPrometheus("echo", nil)
    p.Use(e)
	e.Use(middleware.Recover())
	e.Use(otelecho.Middleware("demo-api"))

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
