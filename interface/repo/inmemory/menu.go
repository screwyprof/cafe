package inmemory

import (
	"github.com/screwyprof/cafe/entity"
)

// MenuRepo stores/retrieves menu items.
type MenuRepo struct {
	menu map[uint8]entity.Item
}

// NewMenuRepo returns a new instance of MenuRepo.
func NewMenuRepo() *MenuRepo {
	return &MenuRepo{
		menu: map[uint8]entity.Item{
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

// All retrieves all menu items.
func (r *MenuRepo) All() map[uint8]entity.Item {
	return r.menu
}
