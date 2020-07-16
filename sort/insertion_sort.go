package sort

// InsertionSort sorts sequence in place in ascending manner
// using insertion sort algorithm
func InsertionSort(sequence []int) {
	for i := 1; i < len(sequence); i++ {
		key := sequence[i]
		j := i - 1
		for j >= 0 && sequence[j] > key {
			sequence[j+1] = sequence[j]
			j--
		}
		sequence[j+1] = key
	}
}
