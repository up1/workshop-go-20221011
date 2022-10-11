//go:build integration

package users_test

import (
	"api/users"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSuccessWithGet(t *testing.T) {
	// Start server
	e := echo.New()
	service := users.NewUserService()
	e.GET("/users", users.GetUserHandler(service))

	rec := httptest.NewRecorder()
	httptest.NewRequest(http.MethodGet, "/users", nil)

	// Assert
	assert.Equal(t, 201, rec.Code)
}
