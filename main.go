package main

import (
	"fmt"

	q "github.com/girishg4t/TDD-samples/lib"
)

func main() {
	fmt.Println("test")
}

func reverse(queue q.Queue) q.Queue {
	rQueue := &q.Queue{}
	rQueue.New()
	s := queue.Size()
	for i := 0; i < s; i++ {

		for j := 0; j < queue.Size()-1; j++ {
			x, _ := queue.Front()
			queue.Dequeue()
			queue.Enqueue(x)
		}

		// Get the last element and
		// add it to the new queue
		el, _ := queue.Front()
		rQueue.Enqueue(el)
		queue.Dequeue()
	}
	return *rQueue
}
