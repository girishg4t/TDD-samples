package queue

import (
	"testing"
)

func TestCheckEmpty(t *testing.T) {
	q := &Queue{}
	resultIsEmpty := q.IsEmpty()

	if !resultIsEmpty {
		t.Errorf("Queue is not empty")
	}

	q.New()
	q.Enqueue(1)
	resultIsNotEmpty := q.IsEmpty()

	if resultIsNotEmpty {
		t.Errorf("Queue should not be empty")
	}

}

func TestEnqueue(t *testing.T) {
	q := &Queue{}
	q.New()
	q.Enqueue(10)
	q.Enqueue(12)
	el, _ := q.Front()
	if el != 10 {
		t.Errorf("Queue does not have element as 10")
	}

	el, _ = q.Front()
	if el != 12 {
		t.Errorf("Queue does not have element as 12")
	}

	q.Enqueue(13)
	q.Enqueue(19)
	el, _ = q.Front()
	if el != 13 {
		t.Errorf("Queue does not have element as 13")
	}
}

func TestDequeue(t *testing.T) {
	q := &Queue{}
	q.New()
	q.Enqueue(10)
	q.Enqueue(12)
	el, err := q.Dequeue()
	if el != 10 {
		t.Errorf("Element not found as 10")
	}

	el, _ = q.Dequeue()
	if el != 12 {
		t.Errorf("Element not found as 12")
	}

	Expected := "Max size reached"

	_, err = q.Front()
	if err.Error() != Expected {
		t.Errorf(Expected)
	}
	_, err = q.Dequeue()
	if err.Error() != Expected {
		t.Errorf(Expected)
	}
}
