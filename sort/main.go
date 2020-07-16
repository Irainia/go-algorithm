package sort

import (
	"fmt"
	"math/rand"
	"time"
)

// Main plays all sorting algorithms
func Main() {
	maxSequenceLength := 60000
	smallSequence := []int{3, 5, 1, 7, 7, 9, 10, 3}
	largeSequence := make([]int, maxSequenceLength)
	for i := 0; i < maxSequenceLength; i++ {
		largeSequence[i] = rand.Int()
	}

	fmt.Println("=== Bubble Sort ===")
	playSmallBubbleSort(smallSequence)
	playLargeBubbleSort(largeSequence)

	fmt.Println()

	fmt.Println("=== Selection Sort ===")
	playSmallSelectionSort(smallSequence)
	playLargeSelectionSort(largeSequence)

	fmt.Println()

	fmt.Println("=== Insertion Sort ===")
	playSmallInsertionSort(smallSequence)
	playLargeInsertionSort(largeSequence)
}

func playSmallBubbleSort(sequence []int) {
	fmt.Printf("before: %+v\n", sequence)
	BubbleSort(sequence)
	fmt.Printf("after: %+v\n", sequence)
}

func playLargeBubbleSort(sequence []int) {
	startTime := time.Now()
	BubbleSort(sequence)
	endTime := time.Now()

	fmt.Printf("sort on %d items takes: %v\n", len(sequence), endTime.Sub(startTime))
}

func playSmallSelectionSort(sequence []int) {
	fmt.Printf("before: %+v\n", sequence)
	SelectionSort(sequence)
	fmt.Printf("after: %+v\n", sequence)
}

func playLargeSelectionSort(sequence []int) {
	startTime := time.Now()
	SelectionSort(sequence)
	endTime := time.Now()

	fmt.Printf("sort on %d items takes: %v\n", len(sequence), endTime.Sub(startTime))
}

func playSmallInsertionSort(sequence []int) {
	fmt.Printf("before: %+v\n", sequence)
	InsertionSort(sequence)
	fmt.Printf("after: %+v\n", sequence)
}

func playLargeInsertionSort(sequence []int) {
	startTime := time.Now()
	InsertionSort(sequence)
	endTime := time.Now()

	fmt.Printf("sort on %d items takes: %v\n", len(sequence), endTime.Sub(startTime))
}
