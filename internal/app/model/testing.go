package model

import "testing"

// Test User
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "example@user.com",
		Password: "password",
	}
}
