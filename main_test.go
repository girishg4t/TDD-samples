package main

import (
	"testing"

	q "github.com/girishg4t/TDD-samples/lib"
)

func TestReverseQueue(t *testing.T) {
	queue := &q.Queue{}
	queue.New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	result := reverse(*queue)

	if result.Values == nil {
		t.Errorf("Queue should not be empty")
	}

	if result.Values[0] != 5 {
		t.Errorf("Element not found 5")
	}
	if result.Values[4] != 1 {
		t.Errorf("Element not found 1")
	}
}
