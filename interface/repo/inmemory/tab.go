package inmemory

import (
	"fmt"
	"github.com/screwyprof/cafe/entity"
)

type tabRepo struct {
	tabs map[string]entity.Tab
}

func NewTabRepo() *tabRepo {
	return &tabRepo{tabs: make(map[string]entity.Tab)}
}

func (r *tabRepo) ByID(ID string) (entity.Tab, error) {
	if u, ok := r.tabs[ID]; ok {
		return u, nil
	}

	return entity.Tab{}, fmt.Errorf("tab %d is not found", ID)
}

func (r *tabRepo) Store(t entity.Tab) error {
	r.tabs[t.ID] = t
	return nil
}
