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

func TestNewAllocatedLinkedList(t *testing.T) {
	t.Run("length should be the same as sequence", func(t *testing.T) {
		sequence := []int{3, 5, 2, 1, 4}
		linkedList := data.NewAllocatedLinkedList(sequence)

		expectedLength := 5

		actualLength := linkedList.GetLength()

		if expectedLength != actualLength {
			t.Errorf("expected: %d -- actual: %d", expectedLength, actualLength)
		}
	})
}
