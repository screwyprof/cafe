package command

type OpenTab struct {
	ID          string
	TableNumber uint8
}

func (c OpenTab) CommandID() string {
	return "OpenTab"
}
