package queue

import (
	"fmt"
	"testing"

	"github.com/deepak-muley/golangexamples/data-structures/queue"
)

func Test_int_queue(t *testing.T) {
	q := new(queue.Queue)
	q.EnqueueInt(10)
	q.EnqueueInt(20)
	q.EnqueueInt(30)
	item, err := q.DequeueInt()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.DequeueInt()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.DequeueInt()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.DequeueInt()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)
}

func Test_string_queue(t *testing.T) {
	q := new(queue.Queue)
	q.EnqueueGeneric("10")
	q.EnqueueGeneric("20")
	q.EnqueueGeneric("30")
	item, err := q.DequeueGeneric()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.DequeueGeneric()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.DequeueGeneric()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)

	item, err = q.DequeueGeneric()
	if err != nil {
		fmt.Println("Failed to deque")
	}
	fmt.Println(item, err)
}
