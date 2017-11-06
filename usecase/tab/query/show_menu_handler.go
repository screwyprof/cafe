package query

import (
	"fmt"
	"github.com/screwyprof/cafe/entity"
	"github.com/screwyprof/cafe/usecase/intf"
)

type showMenuRepository interface {
	All() map[uint8]entity.Item
}

// ShowMenuHandler retrieves menu items.
type ShowMenuHandler struct {
	// TODO: move interfaces to domain repo entity.showMenuRepository
	repo showMenuRepository
}

// NewShowMenuHandler create a new instance of ShowMenuHandler
func NewShowMenuHandler(r showMenuRepository) *ShowMenuHandler {
	return &ShowMenuHandler{repo: r}
}

// Handle handles the given query and returns result.
// Retrieves menu items.
func (h *ShowMenuHandler) Handle(q intf.Query, result interface{}) error {

	r, ok := result.(*ShowMenuResult)
	if !ok {
		return fmt.Errorf("invalid result given")
	}

	var (
		item  MenuItem
		items []MenuItem
	)

	for _, el := range h.repo.All() {
		item = MenuItem{
			Number:      el.MenuNumber,
			Description: el.Description,
			Price:       el.Price,
		}
		items = append(items, item)
	}

	r.Items = items

	return nil
}
