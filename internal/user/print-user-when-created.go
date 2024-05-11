package user

import (
	"errors"
	"fmt"
)

// PrintUserWhenCreated is a handler that creates users
type PrintUserWhenCreated struct{}

func NewPrintUserWhenCreated() *PrintUserWhenCreated {
	return &PrintUserWhenCreated{}
}

func (h *PrintUserWhenCreated) Handle(user interface{}) error {
	u, ok := user.(*User)
	if !ok {
		return errors.New("invalid event received in the handler")
	}

	fmt.Printf("Created user with ID %s, name %s and birthdate %s\n", u.Id(), u.Name(), u.Birthdate().Local())
	return nil
}
