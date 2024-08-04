package user

import (
	"errors"
	"fmt"

	event_bus "github.com/AntonioMartinezFernandez/go-event-bus-example/pkg/event-bus"
)

// SendEmailToUserOnUserCreated is a handler that send emails
type SendEmailToUserOnUserCreated struct{}

func NewSendEmailToUserOnUserCreated() *SendEmailToUserOnUserCreated {
	return &SendEmailToUserOnUserCreated{}
}

func (h *SendEmailToUserOnUserCreated) Handle(event event_bus.Event) error {
	u, ok := event.(*UserCreatedEvent)
	if !ok {
		return errors.New("invalid event received in the handler")
	}

	fmt.Printf("Sending email to user with ID %s, name %s and birthdate %s\n", u.UserId(), u.UserName(), u.UserBirthday().Local())
	return nil
}
