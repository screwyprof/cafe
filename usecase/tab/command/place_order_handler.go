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

	c, ok := cmd.(PlaceOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	tab, err := h.repo.ByID(c.ID)
	if err != nil {
		return fmt.Errorf("tab %q is not opened", c.ID)
	}

	var (
		item entity.Item
	)

	for _, el := range c.Items {
		item = entity.Item{
			MenuNumber:  el.MenuNumber,
			Description: el.Description,
			Price:       el.Price,
			Available:   el.Available,
		}

		if err := tab.Add(item); err != nil {
			return err
		}
	}

	return h.repo.Store(tab)
}
