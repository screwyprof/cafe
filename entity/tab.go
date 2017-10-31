package entity

type Tab struct {
	ID          string
	TableNumber uint8
	Items       []OrderedItem
}
