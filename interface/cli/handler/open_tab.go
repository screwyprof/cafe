package handler

import (
	"github.com/screwyprof/cafe/entity"
	"github.com/screwyprof/cafe/usecase/intf"
	"github.com/screwyprof/cafe/usecase/tab/query"
)

type waiterService interface {
	OpenTab(ID string, tableNumber uint8) error
	PlaceOrder(ID string, items []entity.OrderedItem) error
}

// Waiter manages orders.
type Waiter struct {
	w waiterService
	m intf.QueryHandler
}

// NewWaiter creates a new instance of Waiter
func NewWaiter(w waiterService, m intf.QueryHandler) Waiter {
	return Waiter{w: w, m: m}
}

// OpenTab opens a new Tab.
func (h Waiter) OpenTab(ID string, tableNumber uint8) error {
	return h.w.OpenTab(ID, tableNumber)
}

// PlaceOrder places order with the given orderedItems for the given TabID.
func (h Waiter) PlaceOrder(TabID string, orderedItems []uint8) error {

	menu := &query.ShowMenuResult{}
	h.m.Handle(nil, menu)

	var (
		item  entity.OrderedItem
		items []entity.OrderedItem
	)

	for _, el := range menu.Items {
		item = entity.OrderedItem{
			MenuNumber:  el.Number,
			Description: el.Description,
			Price:       el.Price,
		}
		items = append(items, item)
	}

	return h.w.PlaceOrder(TabID, items)
}
