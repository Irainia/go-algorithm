package sort

import (
	"fmt"
	"math/rand"
	"time"
)

const maxSequenceLength = 100000

var largeSequence []int

func init() {
	getSmallSequence()
	getLargeSequence()
}

// Main plays all sorting algorithms
func Main() {
	fmt.Println("=== Bubble Sort ===")
	playSmallSort(getSmallSequence(), BubbleSort)
	playLargeSort(getLargeSequence(), BubbleSort)

	fmt.Println()

	fmt.Println("=== Selection Sort ===")
	playSmallSort(getSmallSequence(), SelectionSort)
	playLargeSort(getLargeSequence(), SelectionSort)

	fmt.Println()

	fmt.Println("=== Insertion Sort ===")
	playSmallSort(getSmallSequence(), InsertionSort)
	playLargeSort(getLargeSequence(), InsertionSort)

	fmt.Println()

	fmt.Println("=== Merge Sort ===")
	playSmallSort(getSmallSequence(), MergeSort)
	playLargeSort(getLargeSequence(), MergeSort)
}

func getSmallSequence() []int {
	return []int{3, 5, 1, 7, 8, 7, 9, 10, 3}
}

func getLargeSequence() []int {
	largeSeq := make([]int, maxSequenceLength)
	if largeSequence == nil {
		for i := 0; i < maxSequenceLength; i++ {
			largeSeq[i] = rand.Int()
		}
	}
	copy(largeSeq, largeSequence)
	return largeSeq
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
