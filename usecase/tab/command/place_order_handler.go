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

// PlaceOrderHandler stores orders.
type PlaceOrderHandler struct {
	repo placeOrderRepository
}

// NewPlaceOrderHandler creates a new instance of PlaceOrderHandler.
func NewPlaceOrderHandler(r placeOrderRepository) PlaceOrderHandler {
	return PlaceOrderHandler{repo: r}
}

// Handle handles the given command.
// Stores the given order.
func (h PlaceOrderHandler) Handle(cmd intf.Command) error {

	c := cmd.(PlaceOrder)

	tab, err := h.repo.ByID(c.ID)
	if err != nil {
		return fmt.Errorf("tab %q is not opened", c.ID)
	}
	tab.Items = c.Items

	return h.repo.Store(tab)
}
