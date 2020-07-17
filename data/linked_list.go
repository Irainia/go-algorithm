package data

// LinkedList is the type to hold the list in a linked-list manner
type LinkedList struct {
	value int

	nextLinkedList *LinkedList
}

// GetLength gets the length of the current list
func (l *LinkedList) GetLength() int {
	return 0
}

// NewEmptyLinkedList create empty linked-list
func NewEmptyLinkedList() *LinkedList {
	return &LinkedList{}
}
