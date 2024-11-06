package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/microservices/microUserAuth/internal/interface/handlers"
	"github.com/microservices/microUserAuth/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerLogin(t *testing.T) {
	gin.SetMode(gin.TestMode) // ponemos gin en modo de prueba
	mockAuthService := &mocks.MockAuthService{}
	authHandler := handlers.NewAuthHandler(mockAuthService)

	body := `{"username": "testuser", "password": "testpass"}`
	w, c := getContextLogin("POST", "/login", body)
	authHandler.Login(c)

	require.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"token": "mocked-jwt-token"}`, w.Body.String())
}

func getContextLogin(method, url, body string) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(method, url, strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	return w, c
}
