package handler

import (
	"github.com/screwyprof/cafe/usecase/intf"
	"github.com/screwyprof/cafe/usecase/tab/query"
	"github.com/screwyprof/cafe/entity"
)

type waiterService interface {
	OpenTab(ID string, tableNumber uint8) error
	PlaceOrder(ID string, items []entity.OrderedItem) error
}

type Waiter struct {
	w waiterService
	m intf.QueryHandler
}

func NewWaiter(w waiterService, m intf.QueryHandler) Waiter {
	return Waiter{w: w, m: m}
}

func (h Waiter) OpenTab(ID string, tableNumber uint8) error {
	return h.w.OpenTab(ID, tableNumber)
}

func (h Waiter) PlaceOrder(ID string, orderedItems []uint8) error {

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

	return h.w.PlaceOrder(ID, items)
}
