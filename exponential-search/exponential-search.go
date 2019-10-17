package exponential_search

import (
	binary_search "Golang-Algorithms/binary-search"
	"math"
)

func ExponentialSearch(slice []int, key int) int{
	if slice[0] == key{
		return 0
	}

	bound  := 1
	for bound < len(slice) && slice[bound] <= key{
		bound = bound *2
	}

	_, index := binary_search.BinarySearchInt(slice,int(bound / 2), int(math.Min(float64(bound + 1), float64(len(slice) - 1)) ), key)
	return index
}
