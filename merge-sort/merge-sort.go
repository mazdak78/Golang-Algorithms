package merge_sort

type Customer struct {
	Balance int
	ID 	int
	Name string
}

var ToBeSorted []Customer = []Customer{
	Customer{Balance : 20, ID : 1, Name : "John"},
	Customer{Balance : 24, ID : 2, Name: "Joe"},
	Customer{Balance : 122, ID : 3, Name: "Bob"},
	Customer{Balance : 4, ID : 4, Name: "David"},
	Customer{Balance : 200, ID : 5, Name: "Susan"},
	Customer{Balance : 45, ID : 6, Name: "Maria"},
	Customer{Balance : 34, ID : 7, Name: "Ahmad"},
	Customer{Balance : 989, ID : 8, Name: "Katty"},
	Customer{Balance : 1000, ID : 9, Name: "Anai"},
	Customer{Balance : 1, ID : 10, Name: " Ali"},
}

// taken from https://austingwalters.com/merge-sort-in-go-golang/
// Runs MergeSort algorithm on a slice single
func MergeSort(slice []Customer) []Customer {
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func Merge(left, right []Customer) []Customer {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]Customer, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i].Balance < right[j].Balance {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
