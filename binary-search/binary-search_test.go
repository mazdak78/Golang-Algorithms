package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	n := len(ToBeSearched)

	var item, index = BinarySearch(ToBeSearched, 0, n-1, 22)

	if index == -1{
		t.Errorf("Could not find customer with specified age")
	} else{
		fmt.Printf("Found Customer with name %v at index %v \n", item.Name, index)
	}


}
