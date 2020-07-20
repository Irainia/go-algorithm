package example

import (
	"fmt"

	"github.com/Irainia/go-algorithm/data"
)

// PlayData plays all data structures
func PlayData() {
	fmt.Println("\n=== PLAY DATA ===")

	playQueue()

	fmt.Println("=== FINISH ===")
	fmt.Println()
}

func playQueue() {
	listOfValues := []int{3, 5, 6, 1, 2, 1, 4, 7}

	queue := data.Queue{}

	fmt.Println("... Enqueue Process ...")
	for _, value := range listOfValues {
		queue.Enqueue(value)
		fmt.Printf("enqueued: %d -- length: %d\n", value, queue.Length())
	}

	fmt.Println("\n... Dequeu Process ...")
	for queue.Length() > 0 {
		fmt.Printf("dequeued: %d -- length: %d\n", queue.Dequeue(), queue.Length())
	}
}
