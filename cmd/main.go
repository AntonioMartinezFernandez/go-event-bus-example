package main

import (
	"fmt"
	"time"

	user "github.com/AntonioMartinezFernandez/go-event-bus-example/internal/user"
	event_bus "github.com/AntonioMartinezFernandez/go-event-bus-example/pkg/event-bus"
)

func main() {
	// Create a new bus
	bus := event_bus.NewEventBus()

	// Create channel for subscribing to topics
	createdUserChan := make(chan interface{})
	anotherCreatedUserChan := make(chan interface{})

	// Subscribe both channels to a topic
	bus.Subscribe(user.UserCreatedEventName, createdUserChan)
	bus.Subscribe(user.UserCreatedEventName, anotherCreatedUserChan)

	// Handler
	handler := user.NewPrintUserWhenCreated()

	// Run a goroutine to listen for events on the user channel
	go func() {
		for msg := range createdUserChan {
			err := handler.Handle(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	// Run a goroutine to listen for events on the user channel
	go func() {
		for msg := range anotherCreatedUserChan {
			err := handler.Handle(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	// Create an event
	mick := user.NewUser("1", "Mick", time.Now())
	ev, err := user.NewUserCreatedEvent(mick)
	if err != nil {
		panic("error creating user created event")
	}

	// Publish an event
	bus.Publish(ev)

	<-time.After(2000 * time.Millisecond)
}
