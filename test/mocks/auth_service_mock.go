package mocks

import "errors"

type MockAuthService struct{}

func (m *MockAuthService) Login(username, password string) (string, error) {
	if username == "testuser" && password == "testpass" {
		return "mocked-jwt-token", nil
	}
	return "", errors.New("invalid credentials")
}
