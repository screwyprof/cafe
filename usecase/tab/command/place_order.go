package command

import "github.com/screwyprof/cafe/entity"

type PlaceOrder struct {
	ID    string
	Items []entity.OrderedItem
}

func (c PlaceOrder) CommandID() string {
	return "PlaceOrder"
}
