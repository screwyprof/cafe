package query

type MenuItem struct {
	Number      int
	Description string
	Price       float64
}

type ShowMenuResult struct {
	Items []MenuItem
}
