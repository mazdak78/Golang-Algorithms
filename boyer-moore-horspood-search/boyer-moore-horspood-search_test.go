package boyer_moore_horspood_search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	indexes := Search("Hello, we want to find word Guru, so this #phrase has the word Guru.", "word Guru")

	fmt.Printf("Occurance: %v\n", len(indexes))
	for _, i := range indexes  {
		fmt.Printf("Index: %v\n", i)
	}

}
