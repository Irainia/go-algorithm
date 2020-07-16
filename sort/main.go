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
	playSmallSort(smallSequence, BubbleSort)
	playLargeSort(largeSequence, BubbleSort)

	fmt.Println()

	fmt.Println("=== Selection Sort ===")
	playSmallSort(smallSequence, SelectionSort)
	playLargeSort(largeSequence, SelectionSort)

	fmt.Println()

	fmt.Println("=== Insertion Sort ===")
	playSmallSort(smallSequence, InsertionSort)
	playLargeSort(largeSequence, InsertionSort)
}

func playSmallSort(sequence []int, sortAlgorithm func([]int)) {
	fmt.Printf("before: %+v\n", sequence)
	sortAlgorithm(sequence)
	fmt.Printf("after: %+v\n", sequence)
}

func playLargeSort(sequence []int, sortAlgorithm func([]int)) {
	startTime := time.Now()
	sortAlgorithm(sequence)
	endTime := time.Now()

	fmt.Printf("sort on %d items takes: %v\n", len(sequence), endTime.Sub(startTime))
}
