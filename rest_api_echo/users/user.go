package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	Message string `json:"message"`
}

func GetUserHandler(service *UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := UserResponse{service.GetAll()}
		return c.JSON(http.StatusOK, r)
	}
}

// services -> repository
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetAll() string {
	return "Call get user"
}
