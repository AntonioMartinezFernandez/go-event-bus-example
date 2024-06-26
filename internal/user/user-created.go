package user

import (
	"errors"
)

const UserCreatedEventName string = "UserCreatedEvent"

// UserCreatedEvent is an event published when a user is created
type UserCreatedEvent struct {
	data *User
}

func NewUserCreatedEvent(user interface{}) (*UserCreatedEvent, error) {
	u, ok := user.(*User)
	if !ok {
		return nil, errors.New("trying to create user created event with invalid user")
	}

	return &UserCreatedEvent{data: u}, nil
}

func (c *UserCreatedEvent) EventId() string {
	return UserCreatedEventName
}

func (c *UserCreatedEvent) Data() interface{} {
	return c.data
}
