package handler

import (
	"fmt"
	"github.com/screwyprof/cafe/usecase/intf"
	"github.com/screwyprof/cafe/usecase/tab/command"
	"github.com/screwyprof/cafe/usecase/tab/query"
)

type tabInteractor interface {
	OpenTab(c command.OpenTab) error
	PlaceOrder(c command.PlaceOrder) error
}

// Waiter manages orders.
type Waiter struct {
	w tabInteractor
	m intf.QueryHandler
}

// NewWaiter creates a new instance of Waiter
func NewWaiter(w tabInteractor, m intf.QueryHandler) Waiter {
	return Waiter{w: w, m: m}
}

// OpenTab opens a new Tab.
func (h Waiter) OpenTab(ID string, tableNumber uint8) error {
	return h.w.OpenTab(command.OpenTab{ID: ID, TableNumber: tableNumber})
}

// PlaceOrder places order with the given orderedItems for the given TabID.
func (h Waiter) PlaceOrder(TabID string, orderedItems []uint8) error {

	menu, err := h.getMenu()
	if err != nil {
		return err
	}

	//if err := h.validatePlaceOrderRequest(TabID, orderedItems, menu); err != nil {
	//	return err
	//}

	return h.w.PlaceOrder(h.placeOrderCommand(TabID, orderedItems, menu))
}

func (h Waiter) placeOrderCommand(
	TabID string, orderedItems []uint8, menu map[uint8]query.MenuItem) command.PlaceOrder {

	var (
		item  command.OrderedItem
		items []command.OrderedItem
	)

	for _, num := range orderedItems {

		item = command.OrderedItem{}
		item.Available = false
		item.MenuNumber = num

		if menuItem, ok := menu[num]; ok {
			item.Description = menuItem.Description
			item.Price = menuItem.Price
			item.Available = true
		}

		items = append(items, item)
	}

	return command.PlaceOrder{ID: TabID, Items: items}
}

func (h Waiter) validatePlaceOrderRequest(TabID string, orderedItems []uint8, menu map[uint8]query.MenuItem) error {

	for _, idx := range orderedItems {
		if _, ok := menu[idx]; !ok {
			return fmt.Errorf("menu item #%d is not available", idx)
		}
	}

	return nil
}

func (h Waiter) getMenu() (map[uint8]query.MenuItem, error) {

	res := &query.ShowMenuResult{}
	if err := h.m.Handle(nil, res); err != nil {
		return nil, err
	}

	menu := make(map[uint8]query.MenuItem)
	for _, el := range res.Items {
		menu[uint8(el.Number)] = el
	}

	return menu, nil
}
