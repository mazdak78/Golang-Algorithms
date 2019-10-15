package main

import (
	"Golang-Algorithms/bubble-sort"
	merge_sort "Golang-Algorithms/merge-sort"
	"fmt"
)

func main(){


    sorted_bubble := bubble_sort.BubbleSort(bubble_sort.ToBeSorted)

	fmt.Printf("Bubble Sort: \n")
    for i := 0; i < len(sorted_bubble); i++  {
		fmt.Printf("Customer: %s, Balance: %v\n", sorted_bubble[i].Name, sorted_bubble[i].Balance)
	}


	fmt.Printf("---------------------:\n")

	fmt.Printf("Merge Sort: \n")
	sorted_merge := merge_sort.MergeSort(merge_sort.ToBeSorted)

	for i := 0; i < len(sorted_merge); i++  {
		fmt.Printf("Customer: %s, Balance: %v\n", sorted_merge[i].Name, sorted_merge[i].Balance)
	}
}
