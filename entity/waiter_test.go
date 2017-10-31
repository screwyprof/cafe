package entity

import (
	"reflect"
	"testing"
)

func TestOpenTab(t *testing.T) {

	// arrange
	expected := Tab{ID: "Client 1", TableNumber: 7}

	// act
	obtained := Waiter{}.OpenTab("Client 1", 7)

	// assert
	if !reflect.DeepEqual(obtained, expected) {
		t.Errorf("expected %+v, got %+v", expected, obtained)
	}
}
