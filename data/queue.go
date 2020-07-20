package data

// EmptyQueueErrMessage is message given when
// there is illegal call caused by empty queue
const EmptyQueueErrMessage = "queue is empty"

// Queue is type to process queue manner
type Queue struct {
	head *LinkedList
	tail *LinkedList
}

// Peek returns the head value without removing it
func (q *Queue) Peek() int {
	if q.head == nil {
		panic(EmptyQueueErrMessage)
	}
	return q.head.Value
}

// Enqueue enqueues value into the queue
func (q *Queue) Enqueue(value int) {
	linkedList := &LinkedList{
		Value: value,
	}
	q.head = linkedList
	q.tail = linkedList
}
