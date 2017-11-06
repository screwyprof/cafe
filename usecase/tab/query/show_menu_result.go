package query

// MenuItem is a menu item.
type MenuItem struct {
	Number      uint8
	Description string
	Price       float64
}

// ShowMenuResult is query result DTO.
type ShowMenuResult struct {
	Items []MenuItem
}
