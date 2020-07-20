package data_test

import (
	"testing"

	"github.com/Irainia/go-algorithm/data"
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
}
