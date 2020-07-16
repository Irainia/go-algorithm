package sort

// BubbleSort sorts sequence in place in ascending manner
// using bubble sort algorithm
func BubbleSort(sequence []int) {
	for {
		var isSwapped bool
		for i := 0; i < len(sequence)-1; i++ {
			if sequence[i] > sequence[i+1] {
				sequence[i], sequence[i+1] = sequence[i+1], sequence[i]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
}
