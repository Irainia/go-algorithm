package data

// EmptyStackErrMessage is message given when
// there is illegal call caused by empty stack
const EmptyStackErrMessage = "stack is empty"

// Stack is type to process stack manner
type Stack struct {
}

// Peek returns the top value of stack without removing it
func (s *Stack) Peek() int {
	panic(EmptyStackErrMessage)
}
