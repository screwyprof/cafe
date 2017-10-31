package entity

// Waiter manages orders.
type Waiter struct{}

// OpenTab opens a new tab and assigns it to the table
func (w Waiter) OpenTab(ID string, tableNumber uint8) Tab {

	tab := Tab{
		ID:          ID,
		TableNumber: tableNumber,
	}

	return tab
}
