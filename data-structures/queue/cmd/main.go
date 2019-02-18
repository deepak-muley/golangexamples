package main

import (
	"fmt"

	"github.com/deepak-muley/golangexamples/data-structures/queue"
)

func main() {
	var q *queue.Queue = new(queue.Queue)
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	item, err := q.Dequeue()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.Dequeue()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.Dequeue()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.Dequeue()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)
}
