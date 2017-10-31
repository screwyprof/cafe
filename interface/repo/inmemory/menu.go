package inmemory

import (
	"github.com/screwyprof/cafe/entity"
)

type MenuRepo struct {
	menu map[uint8]entity.OrderedItem
}

func NewMenuRepo() *MenuRepo {
	return &MenuRepo{
		menu: map[uint8]entity.OrderedItem{
			1: {
				MenuNumber:  1,
				Description: "pizza",
				Price:       7.5,
			},
			2: {
				MenuNumber:  2,
				Description: "pasta",
				Price:       8.5,
			},
		},
	}
}

func (r *MenuRepo) All() map[uint8]entity.OrderedItem {
	return r.menu
}
