package user

import (
	"errors"
	"fmt"
)

// SendEmailToUserOnUserCreated is a handler that send emails
type SendEmailToUserOnUserCreated struct{}

func NewSendEmailToUserOnUserCreated() *SendEmailToUserOnUserCreated {
	return &SendEmailToUserOnUserCreated{}
}

func (h *SendEmailToUserOnUserCreated) Handle(user interface{}) error {
	u, ok := user.(*User)
	if !ok {
		return errors.New("invalid event received in the handler")
	}

	fmt.Printf("Sending email to user with ID %s, name %s and birthdate %s\n", u.Id(), u.Name(), u.Birthdate().Local())
	return nil
}
