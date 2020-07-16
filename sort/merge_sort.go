package sort

// MergeSort sorts sequence in place in ascending manner
// using merge sort algorithm
func MergeSort(sequence []int) {
	if len(sequence) <= 1 {
		return
	}

	middleIndex := len(sequence) / 2
	MergeSort(sequence[:middleIndex])
	MergeSort(sequence[middleIndex:])

	sortedSequence := make([]int, len(sequence))
	sortedIterator := 0
	leftIterator := 0
	rightIterator := middleIndex
	for sortedIterator < len(sequence) {
		if leftIterator < middleIndex && rightIterator < len(sequence) {
			if sequence[leftIterator] > sequence[rightIterator] {
				sortedSequence[sortedIterator] = sequence[rightIterator]
				rightIterator++
			} else {
				sortedSequence[sortedIterator] = sequence[leftIterator]
				leftIterator++
			}
		} else if leftIterator < middleIndex {
			sortedSequence[sortedIterator] = sequence[leftIterator]
			leftIterator++
		} else {
			sortedSequence[sortedIterator] = sequence[rightIterator]
			rightIterator++
		}
		sortedIterator++
	}
	copy(sequence, sortedSequence)
}
