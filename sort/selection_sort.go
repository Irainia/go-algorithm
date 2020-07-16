package sort

// SelectionSort sorts sequence in place in ascending manner
// using selection sort algorithm
func SelectionSort(sequence []int) {
	for i := 0; i < len(sequence); i++ {
		for j := i + 1; j < len(sequence); j++ {
			if sequence[i] > sequence[j] {
				temp := sequence[i]
				sequence[i] = sequence[j]
				sequence[j] = temp
			}
		}
	}
}
