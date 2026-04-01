package main

import "testing"

func TestSendEmail(t *testing.T) {
	// Mock sendMail
	var lastEmail *email
	originalSendMail := sendMail
	sendMail = func(e *email) {
		lastEmail = e
	}
	defer func() { sendMail = originalSendMail }()

	SendEmail(func(b *EmailBuilder) {
		b.From("test@example.com").
			To("recipient@example.com").
			Subject("Test Subject").
			Body("Test Body")
	})

	if lastEmail == nil {
		t.Fatal("sendMail was not called")
	}

	if lastEmail.from != "test@example.com" {
		t.Errorf("Expected from 'test@example.com', got '%s'", lastEmail.from)
	}
	if lastEmail.to != "recipient@example.com" {
		t.Errorf("Expected to 'recipient@example.com', got '%s'", lastEmail.to)
	}
}

func TestEmailBuilder_From_Validation(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid email, but code did not panic")
		}
	}()

	b := EmailBuilder{}
	b.From("invalid-email")
}
