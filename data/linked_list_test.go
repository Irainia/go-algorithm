package data_test

import (
	"testing"

	"github.com/Irainia/go-algorithm/data"
)

func TestNewEmptyLinkedList(t *testing.T) {
	t.Run("length should be zero", func(t *testing.T) {
		linkedList := data.NewEmptyLinkedList()

		expectedLength := 0

		actualLength := linkedList.GetLength()

		if expectedLength != actualLength {
			t.Errorf("expected: %d -- actual: %d", expectedLength, actualLength)
		}
	})
}
