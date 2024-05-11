package main

import (
	"time"

	user "github.com/AntonioMartinezFernandez/go-event-bus-example/internal/user"
	event_bus "github.com/AntonioMartinezFernandez/go-event-bus-example/pkg/event-bus"
)

func main() {
	// Create a new bus
	bus := event_bus.NewEventBus()

	// Subscribe both channels to a topic
	bus.Subscribe(user.UserCreatedEventName, user.NewPrintUserOnUserCreated())
	bus.Subscribe(user.UserCreatedEventName, user.NewSendEmailToUserOnUserCreated())

	// Create events
	user1 := user.NewUser("1", "Amy", time.Now())
	ev1, err1 := user.NewUserCreatedEvent(user1)
	if err1 != nil {
		panic("error creating user created event")
	}

	user2 := user.NewUser("2", "Louis", time.Now())
	ev2, err2 := user.NewUserCreatedEvent(user2)
	if err2 != nil {
		panic("error creating user created event")
	}

	user3 := user.NewUser("3", "Nina", time.Now())
	ev3, err3 := user.NewUserCreatedEvent(user3)
	if err3 != nil {
		panic("error creating user created event")
	}

	user4 := user.NewUser("4", "Miles", time.Now())
	ev4, err4 := user.NewUserCreatedEvent(user4)
	if err4 != nil {
		panic("error creating user created event")
	}

	user5 := user.NewUser("5", "Duke", time.Now())
	ev5, err5 := user.NewUserCreatedEvent(user5)
	if err5 != nil {
		panic("error creating user created event")
	}

	// Publish events
	bus.Publish(ev1)
	bus.Publish(ev2)
	bus.Publish(ev3)
	bus.Publish(ev4)
	bus.Publish(ev5)

	<-time.After(500 * time.Millisecond)
}
