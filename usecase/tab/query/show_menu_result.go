package query

// MenuItem is a menu item.
type MenuItem struct {
	Number      int
	Description string
	Price       float64
}

// ShowMenuResult is query result DTO.
type ShowMenuResult struct {
	Items []MenuItem
}
