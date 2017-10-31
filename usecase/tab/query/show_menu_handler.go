package query

import (
	"fmt"
	"github.com/screwyprof/cafe/entity"
	"github.com/screwyprof/cafe/usecase/intf"
)

type showMenuRepository interface {
	All() map[uint8]entity.OrderedItem
}

type ShowMenuHandler struct {
	//repo entity.UserRepository
	repo showMenuRepository
}

func NewShowMenuHandler(r showMenuRepository) *ShowMenuHandler {
	return &ShowMenuHandler{repo: r}
}

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
