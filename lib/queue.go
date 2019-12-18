package queue

import (
	"errors"
)

//Queue represent the queue
type Queue struct {
	Values []interface{}
	Rear   int
	Start  int
}

func (queue *Queue) IsEmpty() bool {
	if len(queue.Values) > 0 {
		return false
	}
	return true
}

func (queue *Queue) New() {
	queue.Values = make([]interface{}, 0)
}

func (queue *Queue) Size() int {
	return len(queue.Values)
}

func (queue *Queue) Enqueue(el interface{}) error {
	queue.Values = append(queue.Values, el)
	queue.Rear = queue.Size() - 1
	return nil
}

func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("Max size reached")
	}

	el := queue.Values[0]
	queue.Values = queue.Values[1:]
	queue.Rear = queue.Size() - 1
	queue.Start = 0
	return el, nil
}

func (queue *Queue) Front() (interface{}, error) {
	if queue.Start > queue.Rear {
		return nil, errors.New("Max size reached")
	}
	el := queue.Values[queue.Start]
	queue.Start++
	return el, nil
}
