package test

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

//	func (m *MockAuthService) Login(username, password string) (string, error) {
//		fmt.Println("MockAuthService.Login llamado con:", username, password)
//		if username == "testuser" && password == "testpass" {
//			return "mocked-jwt-token", nil
//		}
//		return "", service.ErrInvalidCredentials
//	}
func (m *MockAuthService) Login(username, password string) (string, error) {
	fmt.Println("MockAuthService.Login llamado con:", username, password)
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}
