package insertion_sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	slice := []int{4,3,6,7,8,1,45,243,12,1000,300}

	InsertionSort(slice)

	for i := 0; i < len(slice) ; i++  {
		fmt.Printf("Item %v : %v \n", i, slice[i])
	}


}
