package events

type Publisher interface {
	Subscribe(handler EventHandler, events ...Event)
	Notify(event Event)
}

type EventHandler interface {
	Notify(event Event)
}

type EventPublisher struct {
	Publisher
	handlers map[string][]EventHandler
}

func (publisher *EventPublisher) Subscribe(handler EventHandler, events ...Event) {
	for _, event := range events {
		handlers := publisher.handlers[event.Name()]
		handlers = append(handlers, handler)
		publisher.handlers[event.Name()] = handlers
	}
}

func (publisher *EventPublisher) Notify(event Event) {
	if event.IsAsynchronous() {
		go publisher.notifyOverRange(event)
	} else {
		publisher.notifyOverRange(event)
	}
}

func (publisher *EventPublisher) notifyOverRange(event Event) {
	for _, handler := range publisher.handlers[event.Name()] {
		handler.Notify(event)
	}
}
