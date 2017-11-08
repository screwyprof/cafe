package entity

import (
	"fmt"
)

// Tab represents a bill.
type Tab struct {
	ID          string
	TableNumber uint8
	Items       []Item
}

// OpenTab opens a new Tab and assigns it to the table.
func OpenTab(ID string, tableNumber uint8) Tab {

	tab := Tab{
		ID:          ID,
		TableNumber: tableNumber,
	}

	return tab
}

// Item represents an ordered menu item.
type Item struct {
	MenuNumber  uint8
	Description string
	Price       float64
	Available   bool
}

// Add adds the given item to the tab.
func (t *Tab) Add(item Item) error {

	if !item.Available {
		return fmt.Errorf("cannot add unavailable items to order: %v", item)
	}

	if t.TotalPrice()+item.Price > 150.00 {
		return fmt.Errorf("a tab may not exceed a total value of $150.00")
	}

	t.Items = append(t.Items, item)
	return nil
}

// TotalPrice calculates total order price.
func (t *Tab) TotalPrice() float64 {
	var sum float64
	for i := range t.Items {
		sum += t.Items[i].Price
	}
	return sum
}
