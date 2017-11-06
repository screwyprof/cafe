package command

import "github.com/screwyprof/cafe/entity"

// PlaceOrder is a command to place an order.
type PlaceOrder struct {
	ID    string
	Items []OrderedItem
}

// CommandID implements Command interface.
func (c PlaceOrder) CommandID() string {
	return "PlaceOrder"
}

// Item represents a menu item which was ordered.
// It is a request model object.
type OrderedItem struct {
	entity.Item
}
