package knuth_morris_pratt_search

import "fmt"

// inspired from https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/
func generateTable(patternBytes []byte, table *[]int) {

	patternLength := len(patternBytes)

	// length of the previous longest prefix suffix
	len := 0
	*table = append(*table, 0) //this is always 0
	i := 1

	for i <  patternLength  {
		if patternBytes[i] == patternBytes[len]{
			len++
			*table = append(*table, len)
			i++
		} else{
			if len != 0{
				len = (*table)[len -1]
			} else{
				*table = append(*table, len)
				i++
			}
		}
	}

}


func Search(text string, pattern string){
	patternBytes := []byte(pattern)
	patternLength := len(patternBytes)

	textBytes := []byte(text)
	textLength := len(textBytes)

	var table []int
	generateTable(patternBytes, &table)

	i , j := 0,0 //i is index of text and j is index of pattern
	for i < textLength  {
		if patternBytes[j] == textBytes[i]{
			i++
			j++
		}
		if j == patternLength{
			fmt.Printf("Found pattern at index %v \n", i -j )
			j = table[j - 1]
		} else if i < textLength && patternBytes[j] != textBytes[i]{
			// mismatch after j matches
			// Do not match p.table[0..p.table[j-1]] characters,
			// they will match anyway
			if j!=0{
				j = table[j - 1]
			} else{
				i = i+1
			}
		}
	}
}
