package command

import (
	"github.com/screwyprof/cafe/entity"
)

type ReceiveOrderRepository interface {
	ByID(ID string) (entity.Tab, error)
	Store(entity.Tab) error
}

type ReceiveOrderHandler struct {
	repo ReceiveOrderRepository
}

func NewReceiveOrderHandler(r ReceiveOrderRepository) ReceiveOrderHandler {
	return ReceiveOrderHandler{repo: r}
}

//
// tab.ReceiveOrder(tab.ID, items)

// pre condition - client has taken a seat at the table
// and got a tab id
// tab.OpenTab(tabID, tableNumber) tab
func (h ReceiveOrderHandler) Handle(c PlaceOrder) error {

	// seatID = SeatID::Generate(tableNumber, seatNumber) = //t1s2
	// tab, err = tabRepo.BySeatID(seatID)

	// if err == NOT_FOUND {
	// tabID = generate(tableNumber, seatNumber) // t1s2tab787914
	// tab = tab.OpenTab(tabID, tableNumber, seatNumber)
	//}

	//  tab.PlaceOrder(items []OrderedItem)
	//  tab.Items = items

	//tab, err := h.repo.ByID(c.ID)
	//if err != nil {
	//	return fmt.Errorf("Tab %q is not opened", c.ID)
	//}
	//
	//tab.Items = c.Items
	//return h.repo.Store(tab)
	return nil
}
