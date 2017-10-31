package entity

// Tab represents a bill.
type Tab struct {
	ID          string
	TableNumber uint8
	Items       []OrderedItem
}
