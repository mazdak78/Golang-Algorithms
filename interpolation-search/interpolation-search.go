package interpolation_search


func InterpolationSearch(slice []int, key int) int{
	low , mid, high := 0, -1, len(slice)- 1

	keepSearching := true
	result := -1
	for keepSearching == true{

		if low == high || slice[low] == slice[high]{
			result = -1
			break
		}

		mid = low + ((high - low) / (slice[high] - slice[low])) * (key - slice[low])
		if slice[mid] == key{
			result = mid
			break;
		} else if slice[mid] < key {
			low = mid + 1
		} else if slice[mid] > key{
			high = mid -1
		}
	}

	return result
}
