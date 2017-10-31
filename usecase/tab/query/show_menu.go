package query

type ShowMenu struct{}

func (q ShowMenu) QueryID() string {
	return "ShowMenu"
}
