package entity

import (
	"reflect"
	"testing"
)

func TestOpenTab(t *testing.T) {

	// arrange
	expected := Tab{ID: "Client 1", TableNumber: 7}

	// act
	obtained := OpenTab("Client 1", 7)

	// assert
	if !reflect.DeepEqual(obtained, expected) {
		t.Errorf("expected %+v, got %+v", expected, obtained)
	}
}

func TestAdd(t *testing.T) {
	// arrange
	expected := &Tab{
		Items:[]Item{
			{
				Price:7.5,
				Available:true,
			},
			{
				Price:5.1,
				Available:true,
			},
		},
	}

	// act
	tab := &Tab{}
	tab.Add(Item{Available:true, Price:7.5})
	tab.Add(Item{Available:true, Price:5.1})

	// assert
	if !reflect.DeepEqual(tab, expected) {
		t.Errorf("expected %+v, got %+v", expected, tab)
	}
}

func TestAdd_TabExceedsTotalValue_ErrorReturned(t *testing.T) {
	// arrange
	expected := "a tab may not exceed a total value of $150.00"

	// act
	tab := &Tab{}
	err := tab.Add(Item{Available:true, Price:151})

	// assert
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}

}

func TestTotalPrice(t *testing.T) {
	// arrange
	expected := 12.6

	tab := &Tab{
		Items:[]Item{
			{
				Price:7.5,
			},
			{
				Price:5.1,
			},
		},
	}

	// act
	obtained := tab.TotalPrice()

	// assert
	if obtained != expected {
		t.Errorf("expected %+v, got %+v", expected, obtained)
	}
}
