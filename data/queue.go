package data

// EmptyQueueErrMessage is message given when
// there is illegal call caused by empty queue
const EmptyQueueErrMessage = "queue is empty"

// Queue is type to process queue manner
type Queue struct {
}

// Peek returns the head value without removing it
func (q *Queue) Peek() int {
	panic(EmptyQueueErrMessage)
}
