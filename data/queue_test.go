package data_test

import (
	"testing"

	"github.com/irainia/go-algorithm/data"
)

func TestQueue(t *testing.T) {
	t.Run("Peek", func(t *testing.T) {
		t.Run("should panic if queue is empty", func(t *testing.T) {
			queue := data.Queue{}

			expectedMessage := data.EmptyQueueErrMessage

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

		t.Run("should return the last value after dequeue until before last", func(t *testing.T) {
			queue := data.Queue{}
			queue.Enqueue(1)
			queue.Enqueue(7)
			queue.Enqueue(4)

			expectedValue := 4

			queue.Dequeue()
			queue.Dequeue()
			actualValue := queue.Peek()

			if actualValue != expectedValue {
				t.Errorf("expected: %d -- actual: %d", expectedValue, actualValue)
			}
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

			defer func() {
				if r := recover(); r != nil {
					actualMessage := r.(string)

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

	t.Run("Length", func(t *testing.T) {
		t.Run("should return zero if queue is empty", func(t *testing.T) {
			queue := data.Queue{}

			expectedLength := 0

			actualLength := queue.Length()

			if actualLength != expectedLength {
				t.Errorf("expected: %d -- actual: %d", expectedLength, actualLength)
			}
		})

		t.Run("should return the number of values inside the queue", func(t *testing.T) {
			queue := data.Queue{}
			queue.Enqueue(1)
			queue.Enqueue(7)
			queue.Enqueue(4)
			queue.Dequeue()

			expectedLength := 2

			actualLength := queue.Length()

			if actualLength != expectedLength {
				t.Errorf("expected: %d -- actual: %d", expectedLength, actualLength)
			}
		})
	})
}
