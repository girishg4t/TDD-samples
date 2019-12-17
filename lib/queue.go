package queue

import (
	"errors"
)

type queueType []interface{}

var queue queueType
var rear = -1
var front = -1

func IsEmpty() bool {
	if len(queue) > 0 {
		return false
	}
	return true
}

func New(n int) queueType {
	rear = 0
	front = 0
	queue = make(queueType, n)
	return queue
}

func Size() int {
	return len(queue)
}

func Enqueue(el interface{}) (queueType, error) {
	if rear < Size() {
		queue[rear] = el
		rear++
		return queue, nil
	}

	return nil, errors.New("Max size reached")
}

func Dequeue() (interface{}, error) {
	if Size() <= 0 {
		return nil, errors.New("Max size reached")
	}

	el := queue[0]
	queue = queue[1:rear]
	rear = Size()
	front = 0
	return el, nil
}

func Front() (interface{}, error) {
	if front >= rear {
		return nil, errors.New("Max size reached")
	}
	el := queue[front]
	front++
	return el, nil
}
