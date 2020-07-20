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
	if q.head == nil {
		q.head = linkedList
	} else {
		q.tail.NextLinkedList = linkedList
	}
	q.tail = linkedList
}

// Dequeue removes and return the first value from queue
func (q *Queue) Dequeue() int {
	if q.head == nil {
		panic(EmptyQueueErrMessage)
	}
	output := q.head.Value
	q.head = q.head.NextLinkedList
	if q.head == nil {
		q.tail = nil
	}
	return output
}
