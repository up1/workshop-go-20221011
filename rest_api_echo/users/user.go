package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	Message string `json:"message"`
}

func UserHandler(c echo.Context) error {
	us := UserService{}
	r := UserResponse{us.GetAll()}
	return c.JSON(http.StatusOK, r)
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

