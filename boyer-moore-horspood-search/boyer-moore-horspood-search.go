package boyer_moore_horspood_search

type badMatchTable struct {
	table   [256]int
	pattern []byte
}

func newBadMatchTable(pattern []byte) *badMatchTable {
	b := badMatchTable{
		pattern: pattern,
	}

	b.table = [256]int{}
	b.table = b.generateTable()

	return &b
}

func (b *badMatchTable) generateTable() [256]int {

	for i := 0; i < 256; i++ {
		b.table[i] = len(b.pattern)
	}

	lastPatternByte := len(b.pattern) - 1

	for i := 0; i < lastPatternByte; i++ {
		b.table[int(b.pattern[i])] = lastPatternByte - i
	}

	return b.table
}

func Search(text, pattern []byte, indexes []int) []int {

	indexes = indexes[:0]

	textLength := len(text)
	patternLength := len(pattern)

	if textLength == 0 || patternLength == 0 || patternLength > textLength {
		return indexes
	}

	lastPatternByte := patternLength - 1

	mt := newBadMatchTable(pattern)
	index := 0
	for index <= (textLength - patternLength) {
		for i := lastPatternByte; text[index+i] == pattern[i]; i-- {
			if i == 0 {
				indexes = append(indexes, index)
				break
			}
		}

		index += mt.table[text[index+lastPatternByte]]
	}

	return indexes
}
