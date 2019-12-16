package queue

import (
	"testing"
)

func TestCheckEmpty(t *testing.T) {

	resultIsEmpty := IsEmpty()

	if !resultIsEmpty {
		t.Errorf("Queue is not empty")
	}

	New(5)
	resultIsNotEmpty := IsEmpty()

	if resultIsNotEmpty {
		t.Errorf("Queue should not be empty")
	}

}

func TestNew(t *testing.T) {
	New(3)
	if Size() != 3 {
		t.Errorf("Not create queue with length as 3")
	}

	New(10)
	if Size() != 10 {
		t.Errorf("Not create queue with length as 10")
	}
}

func TestEnqueue(t *testing.T) {
	New(2)
	Enqueue(10)
	Enqueue(12)
	el, err := Front()
	if el != 10 {
		t.Errorf("Queue does not have element as 10")
	}

	el, _ = Front()
	if el != 12 {
		t.Errorf("Queue does not have element as 12")
	}

	Expected := "Max size reached"
	_, err = Enqueue(13)
	if err.Error() != Expected {
		t.Errorf("Error actual = %v, and Expected = %v.", queue, Expected)
	}

	_, err = Enqueue(19)
	if err.Error() != Expected {
		t.Errorf("Error actual = %v, and Expected = %v.", queue, Expected)
	}

	_, err = Front()
	if err.Error() != Expected {
		t.Errorf("Error actual = %v, and Expected = %v.", queue, Expected)
	}
}

func TestDequeue(t *testing.T) {
	New(2)
	Enqueue(10)
	Enqueue(12)
	el, err := Dequeue()
	if el != 10 {
		t.Errorf("Element not found as 10")
	}

	el, _ = Dequeue()
	if el != 12 {
		t.Errorf("Element not found as 12")
	}

	Expected := "Max size reached"

	_, err = Front()
	if err.Error() != Expected {
		t.Errorf(Expected)
	}
	_, err = Dequeue()
	if err.Error() != Expected {
		t.Errorf(Expected)
	}
}
