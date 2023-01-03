package model

import "testing"

// тестовый хелпер, чтобы структуру юзера не создавать каждый раз заново

func TestUser(t *testing.T) *User {
	// возвращает полностью валидного пользователя
	return &User{
		Email:    "try@example.com",
		Password: "1234",
	}
}
