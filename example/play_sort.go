package example

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/irainia/go-algorithm/sort"
)

const maxSequenceLength = 100000

var largeSequence []int

func init() {
	getSmallSequence()
	getLargeSequence()
}

// PlaySort plays all sorting algorithms
func PlaySort() {
	fmt.Println("\n=== PLAY SORT ===")

	fmt.Println("... Bubble Sort ...")
	playSmallSort(getSmallSequence(), sort.BubbleSort)
	playLargeSort(getLargeSequence(), sort.BubbleSort)

	fmt.Println()

	fmt.Println("... Selection Sort ...")
	playSmallSort(getSmallSequence(), sort.SelectionSort)
	playLargeSort(getLargeSequence(), sort.SelectionSort)

	fmt.Println()

	fmt.Println("... Insertion Sort ...")
	playSmallSort(getSmallSequence(), sort.InsertionSort)
	playLargeSort(getLargeSequence(), sort.InsertionSort)

	fmt.Println()

	fmt.Println("... Merge Sort ...")
	playSmallSort(getSmallSequence(), sort.MergeSort)
	playLargeSort(getLargeSequence(), sort.MergeSort)

	fmt.Println("=== FINISH ===")
	fmt.Println()
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
