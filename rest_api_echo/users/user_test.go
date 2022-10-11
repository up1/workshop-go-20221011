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
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	c := e.NewContext(req, rec)

	service := users.NewUserService()
	users.GetUserHandler(service)(c)

	// Assert
	assert.Equal(t, rec.Code, 200)
	assert.Contains(t, rec.Body.String(), "Call get user")
}

func TestSuccessWithGetWithRealServer(t *testing.T) {
	// Setup router
	e := echo.New()
	service := users.NewUserService()
	e.GET("/users", users.GetUserHandler(service))

	// Call API
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	// Start server
	e.ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, rec.Code, 200)
	assert.Contains(t, rec.Body.String(), "Call get user")
}
