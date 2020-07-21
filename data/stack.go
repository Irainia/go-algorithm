package data

// EmptyStackErrMessage is message given when
// there is illegal call caused by empty stack
const EmptyStackErrMessage = "stack is empty"

// Stack is type to process stack manner
type Stack struct {
	length int

	top *LinkedList
}

// Peek returns the top value of stack without removing it
func (s *Stack) Peek() int {
	if s.top == nil {
		panic(EmptyStackErrMessage)
	}
	return s.top.Value
}

// Length returns length of the stack
func (s *Stack) Length() int {
	return s.length
}

// Push adds value into stack
func (s *Stack) Push(value int) {
	linkedList := &LinkedList{
		Value:          value,
		NextLinkedList: s.top,
	}
	s.top = linkedList
	s.length++
}

// Pop returns last value and remove it from stack
func (s *Stack) Pop() int {
	if s.top == nil {
		panic(EmptyStackErrMessage)
	}
	output := s.top.Value
	s.top = s.top.NextLinkedList
	s.length--
	return output
}
