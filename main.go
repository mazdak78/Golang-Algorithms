package main

import (
	"fmt"
	"Golang-Algorithms/bubble-sort"
)

func main(){
 sorted := bubble_sort.BubbleSort(bubble_sort.ToBeSorted)

	for i := 0; i < len(sorted); i++  {
		fmt.Printf("Customer: %s, Balance: %v\n", sorted[i].Name, sorted[i].Balance)
	}
}
