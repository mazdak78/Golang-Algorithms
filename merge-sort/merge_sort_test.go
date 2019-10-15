package merge_sort

import "testing"

func TestMergeSort(t *testing.T) {
	var sortted = MergeSort(ToBeSorted)

	if sortted[9].Balance != 1000{
		t.Errorf("Error: Our first item should be balance with 1000 balance")
	}

	if sortted[0].Balance != 1{
		t.Errorf("Error: Our first item should be balance with 1000 balance")
	}

}
