package queue

import "fmt"

type Queue struct {
	slice []int
}

func (q *Queue) Enqueue(item int) {
	q.slice = append(q.slice, item)
}

func (q *Queue) Dequeue() (item int, err error) {
	if len(q.slice) == 0 {
		return 0, fmt.Errorf("No elements present in queue")
	}
	item = q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return item, err
}
