package quick_sort

var ToBeSorted []int = []int{234,24,-124,2,35,45,56,232,1,102}

// Lomuto partition scheme

func QuickSort(slice []int, low, high int){
	if low < high{
		p := partition(slice, low, high)

		QuickSort(slice, low, p-1)
		QuickSort(slice, p+1, high)
	}
}

func partition(slice []int, low, high int) int {
	pivot := slice[high]
	i := low
	for j := low; j < high; j++  {
			if slice[j] < pivot{
				slice[i], slice[j] = slice[j], slice[i]
				i++
			}
	}

	slice[i], slice[high] = slice[high], slice[i]
	return i
}
