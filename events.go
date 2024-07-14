package ventana

import (
	"encoding/json"
	"fmt"
)

// Event represents some event.
type Event struct {
	Data interface{} `json:"data"`
	Kind string      `json:"kind"`
}

// NewEvent creates a new event of the given kind.
func NewEvent(kind string) *Event {
	return &Event{Kind: kind}
}

// NewEventWithData creates a new event of the given kind and data.
func NewEventWithData(kind string, data interface{}) *Event {
	return &Event{Kind: kind, Data: data}
}

// SendEvent sends the given event to the UI.
func (window *windowImpl) SendEvent(event *Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	window.Eval(fmt.Sprintf("%s(JSON.parse('%s'))", window.uiEventHandlerName, string(data)))
	return nil
}
