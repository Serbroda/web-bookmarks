package events

type Event struct {
	Name    string
	OldData interface{}
	Data    interface{}
}

type EventListener func(event Event)

type EventDispatcher struct {
	listeners map[string][]EventListener
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		listeners: make(map[string][]EventListener),
	}
}

func (ed *EventDispatcher) RegisterListener(eventName string, listener EventListener) {
	ed.listeners[eventName] = append(ed.listeners[eventName], listener)
}

func (ed *EventDispatcher) Dispatch(event Event) {
	if listeners, exists := ed.listeners[event.Name]; exists {
		for _, listener := range listeners {
			listener(event)
		}
	}
}
