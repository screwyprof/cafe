package command

import (
	"fmt"
	"github.com/screwyprof/cafe/entity"
	"github.com/screwyprof/cafe/usecase/intf"
)

type placeOrderRepository interface {
	ByID(ID string) (entity.Tab, error)
	Store(entity.Tab) error
}

type PlaceOrderHandler struct {
	repo placeOrderRepository
}

func NewPlaceOrderHandler(r placeOrderRepository) PlaceOrderHandler {
	return PlaceOrderHandler{repo: r}
}

// Handle - handles the given command
//
// tab.ReceiveOrder(tab.ID, items)
// pre condition - client has taken a seat at the table
// and got a tab id
// tab.OpenTab(tabID, tableNumber) tab
func (h PlaceOrderHandler) Handle(cmd intf.Command) error {

	c := cmd.(PlaceOrder)

	tab, err := h.repo.ByID(c.ID)
	if err != nil {
		return fmt.Errorf("tab %q is not opened", c.ID)
	}
	tab.Items = c.Items

	return h.repo.Store(tab)
}
