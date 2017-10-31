package cqrs

import "fmt"

// Command defines command pattern interface.
type Command interface {
	CommandID() string
}

// Handler defines command handler pattern interface.
type Handler interface {
	Handle(Command) error
}

// CompositeHandler a command bus.
type CompositeHandler struct {
	handlers map[string]Handler
}

// NewCompositeHandler creates a new instance of CompositeHandler.
func NewCompositeHandler() *CompositeHandler {
	return &CompositeHandler{
		handlers: make(map[string]Handler),
	}
}

// Handle dispatches the given Command to its appropriate Handler and handles it.
func (ch *CompositeHandler) Handle(c Command) error {

	h, err := ch.resolveHandler(c.CommandID())
	if err != nil {
		return err
	}

	return h.Handle(c)
}

func (ch *CompositeHandler) resolveHandler(commandID string) (Handler, error) {

	h, ok := ch.handlers[commandID]
	if !ok {
		return nil, fmt.Errorf("handler for %q is not found", commandID)

	}
	return h, nil
}

// RegisterHandler registers a Handler for the given CommandID.
func (ch *CompositeHandler) RegisterHandler(commandID string, h Handler) *CompositeHandler {
	ch.handlers[commandID] = h
	return ch
}
