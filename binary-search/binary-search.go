package binary_search

type Customer struct {
	Age int
	ID 	int
	Name string
}

// Binary search should be by sorted items
var ToBeSearched []Customer = []Customer{
	Customer{Age : 5,  Name : "John"},
	Customer{Age : 10,  Name: "Joe"},
	Customer{Age : 12,  Name: "Bob"},
	Customer{Age : 18,  Name: "David"},
	Customer{Age : 22,  Name: "Anai"},
	Customer{Age : 34,  Name: "Maria"},
	Customer{Age : 46, Name: "Ahmad"},
	Customer{Age : 50, Name: "Katty"},
	Customer{Age : 53, Name: "Susan"},
	Customer{Age : 70, Name: " Ali"},
}

// taken from https://www.dineshkrish.com/binary-search-in-golang/
func BinarySearchObject(customers []Customer, left int, right int, age int) (Customer, int)  {
	if right >= left {

		mid := left + (right - left) / 2;
		if customers[mid].Age == age {
			return customers[mid], mid
		}

		if customers[mid].Age > age {
			return BinarySearchObject(customers, left, mid-1, age)
		}

		return BinarySearchObject(customers, mid + 1, right, age)
	}

	return Customer{}, -1;
}


// function responsible to perform binary search
func BinarySearchInt(numbers []int, left int, right int, item int) (int, int)  {

	if right >= left {
		mid := left + (right - left) / 2;

		if numbers[mid] == item {
			return numbers[mid], mid
		}

		if numbers[mid] > item {
			return BinarySearchInt(numbers, left, mid-1, item)
		}

		return BinarySearchInt(numbers, mid + 1, right, item)
	}

	return -1, -1
}

