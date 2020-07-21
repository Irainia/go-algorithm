package example

import (
	"fmt"

	"github.com/Irainia/go-algorithm/data"
)

// PlayData plays all data structures
func PlayData() {
	fmt.Println("\n=== PLAY DATA ===")

	fmt.Println("... Queue ...")
	playQueue()

	fmt.Println("... Stack ...")
	playStack()

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

func playStack() {
	listOfValues := []int{3, 5, 6, 1, 2, 1, 4, 7}

	stack := data.Stack{}

	fmt.Println("... Push Process ...")
	for _, value := range listOfValues {
		stack.Push(value)
		fmt.Printf("pushed: %d -- length: %d\n", value, stack.Length())
	}

	fmt.Println("\n... Pop Process ...")
	for stack.Length() > 0 {
		fmt.Printf("popped: %d -- length: %d\n", stack.Pop(), stack.Length())
	}
}
