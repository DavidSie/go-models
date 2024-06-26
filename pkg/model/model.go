package model

import (
	mail "github.com/xhit/go-simple-mail/v2"
)

const (
	RequestEmailTopic string = "RequestEmailTopic"
)


// type EmailRequeststore information regarding requested email
type EmailRequest struct {
	Recipients    []string `yaml:"recipients"`
	CcRecipients  []string `yaml:"cc_recipients"`
	BccRecipients []string `yaml:"bcc_recipients"`
	// If Sender cannot be set iw will be ignored
	Sender             string           `yaml:"sender"`
	Title              string           `yaml:"title"`
	Message            string           `yaml:"message"`
	MessageContentType mail.ContentType `yaml:"message_content_type"`
}
