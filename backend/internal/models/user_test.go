package models

import (
	"testing"
)

func TestUserTableName(t *testing.T) {
	user := User{}
	tableName := user.TableName()
	
	if tableName != "user" {
		t.Errorf("Expected table name 'user', got %s", tableName)
	}
}

func TestUserStruct(t *testing.T) {
	user := User{
		UserName: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	if user.UserName != "testuser" {
		t.Errorf("Expected username 'testuser', got %s", user.UserName)
	}
	if user.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got %s", user.Email)
	}
}

// Note: BeforeCreate tests would require GORM transaction mocking
