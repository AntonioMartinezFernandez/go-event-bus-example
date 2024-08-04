package user

import (
	"errors"
	"time"
)

const UserCreatedEventName string = "UserCreatedEvent"

// UserCreatedEvent is an event published when a user is created
type UserCreatedEvent struct {
	id   string
	data map[string]interface{}
}

func NewUserCreatedEvent(user interface{}) (*UserCreatedEvent, error) {
	u, ok := user.(*User)
	if !ok {
		return nil, errors.New("trying to create user created event with invalid user")
	}
	id := "TODO:ULID"
	data := map[string]interface{}{
		"user_id":  u.Id(),
		"name":     u.Name(),
		"birthday": u.Birthdate(),
	}

	return &UserCreatedEvent{id: id, data: data}, nil
}

func (c *UserCreatedEvent) Id() string {
	return c.id
}

func (c *UserCreatedEvent) Name() string {
	return UserCreatedEventName
}

func (c *UserCreatedEvent) Data() map[string]interface{} {
	return c.data
}

func (c *UserCreatedEvent) UserId() string {
	eventData := c.Data()
	val, _ := eventData["user_id"].(string)
	return val
}

func (c *UserCreatedEvent) UserName() string {
	eventData := c.Data()
	val, _ := eventData["name"].(string)
	return val
}

func (c *UserCreatedEvent) UserBirthday() time.Time {
	eventData := c.Data()
	val, _ := eventData["birthday"].(time.Time)
	return val
}
