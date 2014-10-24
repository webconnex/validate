package validate

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	validEmails := []string{
		"email@example.com",
		"first.last@example.com",
		"first+last@example.com",
		"first++last@example.com",
		"first-last@example.com",
		"first--last@example.com",
		"first_last@example.com",
		"first__last@example.com",
		"_______@example.com",
		"email@subdomain.example.com",
		"1234567890@example.com",
		"email@example-one.com",
		"email@example.name",
		"email@example.museum",
		"email@example.co.jp",
	}
	for _, email := range validEmails {
		if !IsValidEmail(email) {
			t.Errorf("%q should be a valid email", email)
		}
	}
	invalidEmails := []string{
		".email@example.com",
		"email.@example.com",
		"\"email\"@example.com",
		"e'mail.@example.com",
		"first..last@example.com",
		"email@example..com",
		"email@.example.com",
		"email@example.com.",
		"email@example.con",
		"email@example",
		".@example.com",
		"@example.com",
		"example.com",
	}
	for _, email := range invalidEmails {
		if IsValidEmail(email) {
			t.Errorf("%q should not be a valid email", email)
		}
	}
}
