package event_bus

import "fmt"

type EventBus struct {
	subscribers map[string][]chan interface{}
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (b *EventBus) Subscribe(topic string, handler EventHandler) {
	ch := make(chan interface{})
	go func() {
		for msg := range ch {
			err := handler.Handle(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	if _, exists := b.subscribers[topic]; !exists {
		b.subscribers[topic] = make([]chan interface{}, 0)
	}

	b.subscribers[topic] = append(b.subscribers[topic], ch)
}

func (b *EventBus) Unsubscribe(topic string, ch chan interface{}) {
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
	if _, exists := b.subscribers[event.EventId()]; !exists {
		return
	}

	for _, ch := range b.subscribers[event.EventId()] {
		ch <- event.Data()
	}
}

type Event interface {
	EventId() string
	Data() interface{}
}

type EventHandler interface {
	Handle(eventData interface{}) error
}
