package data_test

import (
	"testing"

	"github.com/Irainia/go-algorithm/data"
)

func TestQueue(t *testing.T) {
	t.Run("Peek", func(t *testing.T) {
		t.Run("should panic if queue is empty", func(t *testing.T) {
			queue := data.Queue{}

			expectedMessage := data.EmptyQueueErrMessage

			var actualMessage string
			defer func() {
				if r := recover(); r != nil {
					actualMessage = r.(string)

					if actualMessage != expectedMessage {
						t.Errorf("expected: %s -- actual: %s", expectedMessage, actualMessage)
					}
				}
			}()
			queue.Peek()

			t.Error("expected: panic -- actual: not panic")
		})
	})

	t.Run("Enqueue", func(t *testing.T) {
		t.Run("should enqueue value and peek should return the value", func(t *testing.T) {
			queue := data.Queue{}
			value := 12

			expectedValue := value

			queue.Enqueue(value)
			actualValue := queue.Peek()

			if actualValue != expectedValue {
				t.Errorf("expected: %d -- actual: %d", expectedValue, actualValue)
			}
		})
	})

	t.Run("Dequeue", func(t *testing.T) {
		t.Run("should panic if queue is empty", func(t *testing.T) {
			queue := data.Queue{}

			expectedMessage := data.EmptyQueueErrMessage

			var actualMessage string
			defer func() {
				if r := recover(); r != nil {
					actualMessage = r.(string)

					if actualMessage != expectedMessage {
						t.Errorf("expected: %s -- actual: %s", expectedMessage, actualMessage)
					}
				}
			}()
			queue.Dequeue()

			t.Error("expected: panic -- actual: not panic")
		})

		t.Run("should return the first value from queue if it's not empty", func(t *testing.T) {
			value := 12
			queue := data.Queue{}
			queue.Enqueue(value)

			expectedValue := value
			expectedMessage := data.EmptyQueueErrMessage

			actualValue := queue.Dequeue()

			if actualValue != expectedValue {
				t.Errorf("expected: %d -- actual: %d", expectedValue, actualValue)
			}
			defer func() {
				if r := recover(); r != nil {
					actualMessage := r.(string)

					if actualMessage != expectedMessage {
						t.Errorf("expected: %s -- actual: %s", expectedMessage, actualMessage)
					}
				}
			}()
			queue.Peek()

			t.Error("expected: panic -- actual: not panic")
		})
	})
}
