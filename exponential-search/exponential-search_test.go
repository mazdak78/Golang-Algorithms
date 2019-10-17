package exponential_search

import (
	"fmt"
	"testing"
)

func TestExponentialSearch(t *testing.T) {

	slice := []int{1,10,20,34,45,56,67,78,89,93,102,120,125,220,550,770,880,999}
	key := 45
	index := ExponentialSearch(slice, key)
	fmt.Printf("Founded index for %v is %v\n", key, index)

	key = 1343
	index = ExponentialSearch(slice, key)
	fmt.Printf("Founded index for %v is %v\n", key, index)
}
