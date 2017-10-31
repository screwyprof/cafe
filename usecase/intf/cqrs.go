package intf

type Command interface {
	CommandID() string
}

type CommandHandler interface {
	Handle(Command) error
}

type Query interface {
	QueryID() string
}

type QueryHandler interface {
	Handle(Query, interface{}) error
}
