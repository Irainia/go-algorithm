package data_test

import (
	"testing"

	"github.com/irainia/go-algorithm/data"
)

func TestStack(t *testing.T) {
	t.Run("Peek", func(t *testing.T) {
		t.Run("should panic if stack is empty", func(t *testing.T) {
			stack := data.Stack{}

			expectedMessage := data.EmptyStackErrMessage

			defer func() {
				if r := recover(); r != nil {
					actualMessage := r.(string)

					if actualMessage != expectedMessage {
						t.Errorf("expected: %s -- actual: %s", expectedMessage, actualMessage)
					}
				}
			}()
			stack.Peek()
			t.Error("expected: panic -- actual: not panic")
		})
	})

	t.Run("Length", func(t *testing.T) {
		t.Run("should return zero if stack is empty", func(t *testing.T) {
			stack := data.Stack{}

			expectedLength := 0

			actualLength := stack.Length()

			if actualLength != expectedLength {
				t.Errorf("expected: %d -- actual: %d", expectedLength, actualLength)
			}
		})

		t.Run("should return number of values in stack even after push-pop", func(t *testing.T) {
			stack := data.Stack{}
			stack.Push(3)
			stack.Push(7)
			stack.Push(2)
			stack.Push(8)
			stack.Pop()

			expectedLength := 3

			actualLength := stack.Length()

			if actualLength != expectedLength {
				t.Errorf("expected: %d -- actual: %d", expectedLength, actualLength)
			}
		})
	})

	t.Run("Push", func(t *testing.T) {
		t.Run("should add value and peek return the value", func(t *testing.T) {
			value := 12
			stack := data.Stack{}
			stack.Push(value)

			expectedValue := value

			actualValue := stack.Peek()

			if actualValue != expectedValue {
				t.Errorf("expected: %d -- actual: %d", expectedValue, actualValue)
			}
		})
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("should panic if stack is empty", func(t *testing.T) {
			stack := data.Stack{}

			expectedMessage := data.EmptyStackErrMessage

			defer func() {
				if r := recover(); r != nil {
					actualMessage := r.(string)

					if actualMessage != expectedMessage {
						t.Errorf("expected: %s -- actual: %s", expectedMessage, actualMessage)
					}
				}
			}()
			stack.Pop()
			t.Error("expected: panic -- actual: not panic")
		})

		t.Run("should return the top value and remove it from stack shown by length", func(t *testing.T) {
			value := 12
			stack := data.Stack{}
			stack.Push(value)

			expectedValue := value
			expectedLength := 0

			actualValue := stack.Pop()
			actualLength := stack.Length()

			if actualValue != expectedValue {
				t.Errorf("expected: %d -- actual: %d", expectedValue, actualValue)
			}
			if actualLength != expectedLength {
				t.Errorf("expected: %d -- actual: %d", expectedLength, actualLength)
			}
		})
	})
}
