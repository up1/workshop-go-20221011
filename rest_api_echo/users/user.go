package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("demo-api")

type UserResponse struct {
	Message string `json:"message"`
}

func GetUserHandler(service *UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, span := tracer.Start(c.Request().Context(), "GetUserHandler", oteltrace.WithAttributes(attribute.String("layer", "handler")))
		defer span.End()
		r := UserResponse{service.GetAll()}
		return c.JSON(http.StatusOK, r)
	}
}

// services -> repository
type UserService struct {
	repo IRepository
}

func NewUserService(repo IRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetAll() string {
	result, _ := us.repo.GetSth()
	return result
}
