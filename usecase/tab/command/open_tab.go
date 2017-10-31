package command

// OpenTab is a command to open a new tab.
type OpenTab struct {
	ID          string
	TableNumber uint8
}

// CommandID implements Command interface.
func (c OpenTab) CommandID() string {
	return "OpenTab"
}
