package bubble_sort



type Customer struct {
	Balance int
	ID 	int
	Name string
}

var ToBeSorted [10]Customer = [10]Customer{
	Customer{Balance : 20, ID : 1, Name : "John"},
	Customer{Balance : 24, ID : 2, Name: "Joe"},
	Customer{Balance : 122, ID : 3, Name: "Bob"},
	Customer{Balance : 4, ID : 4, Name: "David"},
	Customer{Balance : 200, ID : 5, Name: "Susan"},
	Customer{Balance : 45, ID : 6, Name: "Maria"},
	Customer{Balance : 34, ID : 7, Name: "David"},
	Customer{Balance : 989, ID : 8, Name: "Katty"},
	Customer{Balance : 1000, ID : 9, Name: "Anai"},
	Customer{Balance : 1, ID : 10, Name: " Ali"},
}

// BubbleSort taken from https://tutorialedge.net/golang/implementing-bubble-sort-with-golang/
func BubbleSort(input [10]Customer) [10]Customer{
	// n is the number of items in our list
	n := 10

	// set swapped to true
	swapped := true

	// iterate through all of the elements in our list
	for swapped{
		swapped = false
		for i := 1; i < n; i++  {
			// element, swap them
			if input[i-1].Balance < input[i].Balance {

				// swap values using Go's tuple assignment
				input[i], input[i-1] = input[i-1], input[i]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}

	return input
}
