package query

// ShowMenu is a query to show menu.
type ShowMenu struct{}

// QueryID implements Query interface.
func (q ShowMenu) QueryID() string {
	return "ShowMenu"
}
