package command

import (
	"github.com/screwyprof/cafe/entity"
	"github.com/screwyprof/cafe/usecase/intf"
)

type openTabRepository interface {
	Store(entity.Tab) error
}

type OpenTabHandler struct {
	repo openTabRepository
}

func NewOpenTabHandler(repo openTabRepository) *OpenTabHandler {
	return &OpenTabHandler{repo: repo}
}

func (h *OpenTabHandler) Handle(cmd intf.Command) error {

	c := cmd.(OpenTab)

	w := entity.Waiter{}
	tab := w.OpenTab(c.ID, c.TableNumber)

	return h.repo.Store(tab)
}
