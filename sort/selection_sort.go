package sort

// SelectionSort sorts sequence in place in ascending manner
// using selection sort algorithm
func SelectionSort(sequence []int) {
	for i := 0; i < len(sequence); i++ {
		compareIndex := i
		for j := i + 1; j < len(sequence); j++ {
			if sequence[compareIndex] > sequence[j] {
				compareIndex = j
			}
		}
		if compareIndex != i {
			sequence[i], sequence[compareIndex] = sequence[compareIndex], sequence[i]
		}
	}
}
