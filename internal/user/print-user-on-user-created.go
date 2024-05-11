package user

import (
	"errors"
	"fmt"
)

// PrintUserOnUserCreated is a handler that print users
type PrintUserOnUserCreated struct{}

func NewPrintUserOnUserCreated() *PrintUserOnUserCreated {
	return &PrintUserOnUserCreated{}
}

func (h *PrintUserOnUserCreated) Handle(user interface{}) error {
	u, ok := user.(*User)
	if !ok {
		return errors.New("invalid event received in the handler")
	}

	fmt.Printf("Printing user with ID %s, name %s and birthdate %s\n", u.Id(), u.Name(), u.Birthdate().Local())
	return nil
}
