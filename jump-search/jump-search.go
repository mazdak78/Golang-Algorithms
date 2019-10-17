package jump_search

import (
	"math"
)

func JumpSearch(slice []int, key int) int{
	step := int(math.Sqrt(float64(len(slice))))
	size := len(slice)

	// Finding the block where element is
	// present (if it is present)
	prev := 0

	for slice[int(math.Min(float64(step), float64(size)))-1] < key{
		prev = step
		step += int(math.Sqrt(float64(size)))

		//key was not found
		if prev >= int(size){
			return -1
		}
	}

	// do linear search in the prev block
	for slice[prev] < key{
		prev++
		// If we reached next block or end of
		// array, element is not present.
		if prev == int(math.Min(float64(step), float64(size))){
			return -1;
		}
	}

	// If element is found
	if slice[prev] == key{
		return prev;
	}

	return -1
}
