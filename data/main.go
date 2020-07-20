package data

import "fmt"

// Main plays all data structures
func Main() {
	playQueue()
}

func playQueue() {
	listOfValues := []int{3, 5, 6, 1, 2, 1, 4, 7}

	queue := Queue{}

	fmt.Println("=== Enqueue Process ===")
	for _, value := range listOfValues {
		queue.Enqueue(value)
		fmt.Printf("enqueued: %d -- length: %d\n", value, queue.Length())
	}

	fmt.Println("\n=== Dequeu Process ===")
	for queue.Length() > 0 {
		fmt.Printf("dequeued: %d -- length: %d\n", queue.Dequeue(), queue.length)
	}
}
