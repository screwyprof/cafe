package inmemory

import (
	"fmt"
	"github.com/screwyprof/cafe/entity"
)

// TabRepo stores/retrieves tabs.
type TabRepo struct {
	tabs map[string]entity.Tab
}

// NewTabRepo returns a new instance of TabRepo.
func NewTabRepo() *TabRepo {
	return &TabRepo{tabs: make(map[string]entity.Tab)}
}

// ByID retrieves a Tab by its ID.
func (r *TabRepo) ByID(ID string) (entity.Tab, error) {
	if u, ok := r.tabs[ID]; ok {
		return u, nil
	}

	return entity.Tab{}, fmt.Errorf("tab %q is not found", ID)
}

// Store stores the given Tab.
func (r *TabRepo) Store(t entity.Tab) error {
	r.tabs[t.ID] = t
	return nil
}
