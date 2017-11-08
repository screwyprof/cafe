package command

import (
	"reflect"
	"testing"

	"github.com/screwyprof/cafe/entity"
)

type OpenTabRepoFake struct {
	Tab entity.Tab
}

func (r *OpenTabRepoFake) Store(tab entity.Tab) error {
	r.Tab = tab
	return nil
}

func TestOpenTab_ValidCommandGiven_TabOpened(t *testing.T) {
	// arrange
	expected := entity.Tab{ID: "Client 1", TableNumber: 5}

	// act
	repo := &OpenTabRepoFake{}
	h := NewOpenTabHandler(repo)
	h.Handle(OpenTab{ID: "Client 1", TableNumber: 5})

	// assert
	if !reflect.DeepEqual(expected, repo.Tab) {
		t.Errorf("expected %+v, got %+v", expected, repo.Tab)
	}
}
