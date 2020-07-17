package data

// LinkedList is the type to hold the list in a linked-list manner
type LinkedList struct {
	value int
	count int

	nextLinkedList *LinkedList
}

// GetLength gets the length of the current list
func (l *LinkedList) GetLength() int {
	return l.count
}

// NewEmptyLinkedList create empty linked-list
func NewEmptyLinkedList() *LinkedList {
	return &LinkedList{}
}

// NewAllocatedLinkedList create pre-allocated linked-list based on sequence
func NewAllocatedLinkedList(sequence []int) *LinkedList {
	return &LinkedList{
		count: len(sequence),
	}
}
