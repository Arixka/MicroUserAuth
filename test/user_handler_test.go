package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/microservices/microUserAuth/internal/interface/handlers"
	"github.com/microservices/microUserAuth/internal/usecase/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	gin.SetMode(gin.TestMode)
}
func TestHandlerLogin_Success(t *testing.T) {
	mockAuthService := &MockAuthService{}
	mockAuthService.On("Login", "testuser", "testpass").Return("mocked-jwt-token", nil)
	authHandler := handlers.NewAuthHandler(mockAuthService)

	body := `{"username": "testuser", "password": "testpass"}`
	w, c := getContextLogin("POST", "/login", body)
	authHandler.Login(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"token": "mocked-jwt-token"}`, w.Body.String())
}
func TestHandlerLogin_InvalidCredentials(t *testing.T) {
	mockAuthService := &MockAuthService{}
	mockAuthService.On("Login", "usermal", "passmal").Return("", service.ErrInvalidCredentials)
	authHandler := handlers.NewAuthHandler(mockAuthService)

	body := `{"username": "usermal", "password": "passmal"}`
	w, c := getContextLogin("POST", "/login", body)
	authHandler.Login(c)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "invalid credentials"}`, w.Body.String())
}

func TestHandlerLogin_InternalServerError(t *testing.T) {
	//sobreescribimos el mock para pasarle un error inesperado
	mockAuthService := &MockAuthService{}
	mockAuthService.On("Login", "usererror", "passerror").Return("", service.ErrInternalServerError)

	authHandler := handlers.NewAuthHandler(mockAuthService)

	body := `{"username": "usererror", "password": "passerror"}`
	w, c := getContextLogin("POST", "/login", body)
	authHandler.Login(c)

	mockAuthService.AssertExpectations(t)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "internal server error"}`, w.Body.String())
}

func getContextLogin(method, url, body string) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(method, url, strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	return w, c
}
