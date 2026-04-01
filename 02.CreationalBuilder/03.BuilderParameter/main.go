package main

import "strings"

// email represents an email message.
type email struct {
	from, to, subject, body string
}

// EmailBuilder handles the construction of email messages.
type EmailBuilder struct {
	email email
}

// From sets the sender address. It validates that the email contains '@'.
func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

// To sets the recipient address.
func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

// Subject sets the email subject.
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

// Body sets the email body.
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

// sendMail is the function responsible for sending the email.
// It is a variable to allow mocking in tests.
var sendMail = func(email *email) {
	// actually ends the email
	// in a real app, this would perform network I/O
}

type build func(*EmailBuilder)

// SendEmail constructs an email using the provided action and sends it.
func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMail(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}