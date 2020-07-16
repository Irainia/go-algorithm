package sort

import (
	"fmt"
	"math/rand"
	"time"
)

// Main plays all sorting algorithms
func Main() {
	sequence := []int{3, 5, 1, 7, 7, 9, 10, 3}

	fmt.Println("=== Bubble Sort ===")
	playMiniBubbleSort(sequence)
	playLargeBubbleSort(100000)
}

func playMiniBubbleSort(sequence []int) {
	fmt.Printf("before: %+v\n", sequence)
	BubbleSort(sequence)
	fmt.Printf("after: %+v\n", sequence)
}

func playLargeBubbleSort(maxSequenceLength uint) {
	sequence := make([]int, maxSequenceLength)

	for i := uint(0); i < maxSequenceLength; i++ {
		sequence[i] = rand.Int()
	}

	startTime := time.Now()
	BubbleSort(sequence)
	endTime := time.Now()

	fmt.Printf("sort on %d items takes: %v\n", maxSequenceLength, endTime.Sub(startTime))
}
