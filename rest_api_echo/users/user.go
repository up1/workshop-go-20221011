package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	Message string `json:"message"`
}

func UserHandler(c echo.Context) error {
	r := UserResponse{"Call get user"}
	return c.JSON(http.StatusOK, r)
}