package command

import (
	"github.com/screwyprof/cafe/entity"
	"errors"
	"reflect"
	"strings"
	"testing"
)

type PlaceOrderRepoMock struct {
	Tab entity.Tab
	Err error
}

func (r *PlaceOrderRepoMock) ByID(ID string) (entity.Tab, error) {

	if r.Err != nil {
		return r.Tab, r.Err
	}

	return r.Tab, nil
}

func (r *PlaceOrderRepoMock) Store(tab entity.Tab) error {
	r.Tab = tab
	return nil
}

func TestPlaceOrder_ValidCommandGiven_OrderPlaced(t *testing.T) {

	// arrange
	tab := entity.Tab{
		ID: "Client 1",
	}

	expected := entity.Tab{
		ID: "Client 1",
		Items: []entity.OrderedItem{
			{
				MenuNumber:  1,
				Description: "pizza",
				Price:       8.5,
			},
		},
	}

	// act
	repo := &PlaceOrderRepoMock{Tab: tab}
	h := NewPlaceOrderHandler(repo)

	c := PlaceOrder{
		ID: "Client 1",
		Items: []entity.OrderedItem{
			{
				MenuNumber:  1,
				Description: "pizza",
				Price:       8.5,
			},
		},
	}

	if err := h.Handle(c); err != nil {
		t.Fatal(err.Error())
	}

	// assert
	if !reflect.DeepEqual(expected, repo.Tab) {
		t.Errorf("expected %+v, got %+v", expected, repo.Tab)
	}
}

func TestPlaceOrder_TabIsNotOpened_OrderNotPlaced(t *testing.T) {

	// arrange
	expected := "is not opened"

	// act
	repo := &PlaceOrderRepoMock{Err: errors.New("some error")}
	h := NewPlaceOrderHandler(repo)

	c := PlaceOrder{
		ID: "Client 1",
		Items: []entity.OrderedItem{
			{
				MenuNumber:  1,
				Description: "pizza",
				Price:       8.5,
			},
		},
	}

	err := h.Handle(c)
	if err == nil {
		t.Fatal("an error expected, but got nil")
	}

	if !strings.Contains("is not opened", expected) {
		t.Fatal(err.Error())
	}

}
