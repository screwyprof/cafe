package tab

import (
	"github.com/screwyprof/cafe/usecase/intf"
	"github.com/screwyprof/cafe/usecase/tab/command"
)

// Interactor a facade to Tab.
type Interactor struct {
	commandBus intf.CommandHandler
}

// NewInteractor creates a new instance of Interactor.
func NewInteractor(bus intf.CommandHandler) Interactor {
	return Interactor{commandBus: bus}
}

// OpenTab opens a new Tab for the given table.
func (i Interactor) OpenTab(c command.OpenTab) error {
	return i.commandBus.Handle(c)
}

// PlaceOrder places a new order.
func (i Interactor) PlaceOrder(c command.PlaceOrder) error {
	return i.commandBus.Handle(c)
}

// Order - request dto
//type Order struct {
//	TableNum uint8
//	SeatNum  uint8
//	Items    []uint8
//}
//
//type SeatIDGenerator interface {
//	Generate(tableNum, seatNum uint8) string
//}
//
//type TabIDGenerator interface {
//	Generate(seatID string) string
//}
//
//type TabBySeat interface {
//	IDBySeatID(seatID string) (string, error)
//}
//
//func (i Interactor) ReceiveOrder(o Order,
//	seatIDGenerator SeatIDGenerator,
//	tabs TabBySeat,
//	tabIDGenerator TabIDGenerator) error {
//
//	seatID := seatIDGenerator.Generate(o.TableNum, o.SeatNum)
//
//	var tabID string
//	tabID, err := tabs.IDBySeatID(seatID)
//	if err != nil {
//		tabID = tabIDGenerator.Generate(seatID)
//	}
//
//	return i.PlaceOrder(tabID, o.Items)
//}

//func Init() {

// client.TakeASeat(tableNum, seatNum)

// client.PlaceOrder(orderedItems),
// tab.ReceiveOrder(Order)

// tab.ServeFood(seatID)
// client.Eat()

// client.Pay(tableNum, seatNum)
//   // tab.ReceivePayment(Payment)
// tab.CloseTab(tabID)
// cashier.AcceptPayment()
//}
