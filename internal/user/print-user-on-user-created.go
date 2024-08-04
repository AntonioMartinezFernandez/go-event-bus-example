package user

import (
	"errors"
	"fmt"

	event_bus "github.com/AntonioMartinezFernandez/go-event-bus-example/pkg/event-bus"
)

// PrintUserOnUserCreated is a handler that print users
type PrintUserOnUserCreated struct{}

func NewPrintUserOnUserCreated() *PrintUserOnUserCreated {
	return &PrintUserOnUserCreated{}
}

func (h *PrintUserOnUserCreated) Handle(event event_bus.Event) error {
	u, ok := event.(*UserCreatedEvent)
	if !ok {
		return errors.New("invalid event received in the handler")
	}

	fmt.Printf("Printing user with ID %s, name %s and birthdate %s\n", u.UserId(), u.UserName(), u.UserBirthday().Local())
	return nil
}
