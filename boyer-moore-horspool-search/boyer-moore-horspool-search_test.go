package boyer_moore_horspool_search

import (
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	c := []byte("Hello, we want to find word Guru, so this #phrase has the word Guru.")
	s := []byte("word Guru")
	indexes := make([]int, 0, 100)

	for i := 0; i < b.N; i++ {
		Search(c, s, indexes)
	}
}
