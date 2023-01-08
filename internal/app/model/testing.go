package model

import (
	"testing"
	"time"
)
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
func TestNote(t *testing.T) *Note {
	return &Note{
		Header:    "header",
		Body:      "body",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
