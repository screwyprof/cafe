package tab

import (
	"github.com/screwyprof/cafe/entity"
	"github.com/screwyprof/cafe/usecase/intf"
	"github.com/screwyprof/cafe/usecase/tab/command"
)

// Aggregate a facade to Tab.
type Aggregate struct {
	commandBus intf.CommandHandler
}

// NewUseCase creates a new instance of Aggregate.
func NewUseCase(bus intf.CommandHandler) Aggregate {
	return Aggregate{commandBus: bus}
}

// OpenTab opens a new Tab for the given table.
func (a Aggregate) OpenTab(ID string, tableNumber uint8) error {
	return a.commandBus.Handle(command.OpenTab{ID: ID, TableNumber: tableNumber})
}

// PlaceOrder places a new order.
func (a Aggregate) PlaceOrder(ID string, items []entity.OrderedItem) error {
	return a.commandBus.Handle(command.PlaceOrder{ID: ID, Items: items})
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
//func (a Aggregate) ReceiveOrder(o Order,
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
//	return a.PlaceOrder(tabID, o.Items)
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
