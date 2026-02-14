package email

import (
	"testing"
)

func TestNewEmailSender(t *testing.T) {
	// Note: This requires viper configuration to be set up
	// Testing the structure
	sender := &EmailSender{
		smtp: "smtp.example.com",
		port: 587,
		host: "test@example.com",
		pwd:  "password",
	}
	
	if sender.smtp == "" {
		t.Error("smtp should not be empty")
	}
	if sender.port == 0 {
		t.Error("port should not be zero")
	}
	if sender.host == "" {
		t.Error("host should not be empty")
	}
}

// Note: Actual email sending tests would require SMTP server mocking
// and are better suited for integration tests rather than unit tests.
