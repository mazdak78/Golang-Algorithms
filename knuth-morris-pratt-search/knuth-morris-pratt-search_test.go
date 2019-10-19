package knuth_morris_pratt_search

import (
	"testing"
)

func TestSearch(t *testing.T) {

	Search("bacbababababcaabaabcbab abababca", "abababca")
}
