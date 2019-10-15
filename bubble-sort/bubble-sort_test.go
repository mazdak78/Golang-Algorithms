package bubble_sort

import "testing"

func TestBubbleSort(t *testing.T) {
	var sortted = BubbleSort(ToBeSorted)

	if sortted[0].Balance != 1000{
		t.Errorf("Error: Our first item should be balance with 1000 balance")
	}

	if sortted[9].Balance != 1{
		t.Errorf("Error: Our first item should be balance with 1000 balance")
	}

}
