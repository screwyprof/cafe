package command

import (
	"github.com/screwyprof/cafe/entity"
	"github.com/screwyprof/cafe/usecase/intf"
)

type openTabRepository interface {
	Store(entity.Tab) error
}

// OpenTabHandler creates a new Tab.
type OpenTabHandler struct {
	repo openTabRepository
}

// NewOpenTabHandler creates a new instance of OpenTabHandler.
func NewOpenTabHandler(repo openTabRepository) *OpenTabHandler {
	return &OpenTabHandler{repo: repo}
}

// Handle handles the given command.
// Creates a new Tab.
func (h *OpenTabHandler) Handle(cmd intf.Command) error {
	c := cmd.(OpenTab)
	return h.repo.Store(entity.OpenTab(c.ID, c.TableNumber))
}
