package cqrs

import "fmt"

type Command interface {
	CommandID() string
}

type Handler interface {
	Handle(Command) error
}

type compositeHandler struct {
	handlers map[string]Handler
}

func NewCompositeHandler() *compositeHandler {
	return &compositeHandler{
		handlers: make(map[string]Handler),
	}
}

func (ch *compositeHandler) Handle(c Command) error {

	h, err := ch.resolveHandler(c.CommandID())
	if err != nil {
		return err
	}

	return h.Handle(c)
}

func (ch *compositeHandler) resolveHandler(commandID string) (Handler, error) {

	h, ok := ch.handlers[commandID]
	if !ok {
		return nil, fmt.Errorf("handler for %q is not found", commandID)

	}
	return h, nil
}

func (ch *compositeHandler) RegisterHandler(commandID string, h Handler) *compositeHandler {
	ch.handlers[commandID] = h
	return ch
}
