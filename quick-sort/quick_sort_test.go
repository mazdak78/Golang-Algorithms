package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	QuickSort(ToBeSorted, 0, len(ToBeSorted) -1)

	for i,s := range(ToBeSorted)  {
		fmt.Printf("a[%d] = %d\n", i,s)
	}

}
