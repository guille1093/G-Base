package tests

import (
	"github.com/guille1093/G-Base/tools/mailer"
)

var _ mailer.Mailer = (*TestMailer)(nil)

// TestMailer is a mock `mailer.Mailer` implementation.
type TestMailer struct {
	TotalSend   int
	LastMessage mailer.Message
}

// Reset clears any previously test collected data.
func (m *TestMailer) Reset() {
	m.TotalSend = 0
	m.LastMessage = mailer.Message{}
}

// Send implements `mailer.Mailer` interface.
func (c *TestMailer) Send(m *mailer.Message) error {
	c.TotalSend++
	c.LastMessage = *m

	return nil
}
