package example

import (
	"fmt"

	"github.com/irainia/go-algorithm/data"
	"github.com/irainia/go-algorithm/uf"
)

// PlayData plays all data structures
func PlayData() {
	fmt.Println("\n=== PLAY DATA ===")

	fmt.Println("... Queue ...")
	playQueue()

	fmt.Println("... Stack ...")
	playStack()

	fmt.Println("... Union-Find ...")
	playUnionFind()

	fmt.Println("=== FINISH ===")
	fmt.Println()
}

func playUnionFind() {
	maximumNumber := 10
	unionFind := uf.NewUnionFind(maximumNumber)

	fmt.Println("... First Union ...")
	firstPointNumber := [][]int{
		[]int{4, 3},
		[]int{3, 8},
		[]int{6, 5},
		[]int{9, 4},
		[]int{2, 1},
	}
	for _, points := range firstPointNumber {
		unionFind.Union(points[0], points[1])
		fmt.Printf("union of [%d] and [%d]\n", points[0], points[1])
	}
	fmt.Println("... First Check ...")
	firstCheckPoint := [][]int{
		[]int{0, 7},
		[]int{8, 9},
	}
	for _, points := range firstCheckPoint {
		if unionFind.IsConnected(points[0], points[1]) {
			fmt.Printf("[%d] and [%d] is connected\n", points[0], points[1])
		} else {
			fmt.Printf("[%d] and [%d] is not connected\n", points[0], points[1])
		}
	}
	fmt.Println("... Second Union ...")
	secondPointNumber := [][]int{
		[]int{5, 0},
		[]int{7, 2},
		[]int{6, 1},
		[]int{1, 0},
	}
	for _, points := range secondPointNumber {
		unionFind.Union(points[0], points[1])
		fmt.Printf("union of [%d] and [%d]\n", points[0], points[1])
	}
	fmt.Println("... Second Check ...")
	secondCheckPoint := [][]int{
		[]int{0, 7},
	}
	for _, points := range secondCheckPoint {
		if unionFind.IsConnected(points[0], points[1]) {
			fmt.Printf("[%d] and [%d] is connected\n", points[0], points[1])
		} else {
			fmt.Printf("[%d] and [%d] is not connected\n", points[0], points[1])
		}
	}
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
	fmt.Println()
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
	fmt.Println()
}
