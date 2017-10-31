package command

import "github.com/screwyprof/cafe/entity"

// PlaceOrder is a command to place an order.
type PlaceOrder struct {
	ID    string
	Items []entity.OrderedItem
}

// CommandID implements Command interface.
func (c PlaceOrder) CommandID() string {
	return "PlaceOrder"
}
