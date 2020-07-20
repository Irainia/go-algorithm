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
}
