package insertion_sort

func InsertionSort(slice []int){
	n := len(slice)

	for i := 1; i < n ; i++ {
		key := slice[i]
		j := i

		for j -1 >= 0 && slice[j-1] > key{
			slice[j] = slice[j-1]

			j = j -1
		}

		slice[j] = key
	}
}
