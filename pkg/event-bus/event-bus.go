package event_bus

import "fmt"

type EventBus struct {
	subscribers map[string][]chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan Event),
	}
}

func (b *EventBus) Subscribe(topic string, handler EventHandler) {
	ch := make(chan Event)
	go func() {
		for msg := range ch {
			err := handler.Handle(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	if _, exists := b.subscribers[topic]; !exists {
		b.subscribers[topic] = make([]chan Event, 0)
	}

	b.subscribers[topic] = append(b.subscribers[topic], ch)
}

func (b *EventBus) Unsubscribe(topic string, ch chan Event) {
	if _, exists := b.subscribers[topic]; !exists {
		return
	}

	for i, subscriber := range b.subscribers[topic] {
		if subscriber == ch {
			// Remove the channel from the slice
			b.subscribers[topic] = append(b.subscribers[topic][:i], b.subscribers[topic][i+1:]...)
			break
		}
	}
}

func (b *EventBus) Publish(event Event) {
	if _, exists := b.subscribers[event.Name()]; !exists {
		return
	}

	for _, ch := range b.subscribers[event.Name()] {
		ch <- event
	}
}

type Event interface {
	Name() string
	Id() string
	Data() map[string]interface{}
}

type EventHandler interface {
	Handle(event Event) error
}
