package queue

import "fmt"

type Queue struct {
	slice []interface{}
}

func (q *Queue) EnqueueInt(item int) {
	q.slice = append(q.slice, item)
}

func (q *Queue) EnqueueGeneric(item interface{}) {
	q.slice = append(q.slice, item)
}

func (q *Queue) DequeueInt() (item int, err error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("No elements present in queue")
	}
	item, ok := (q.slice[0]).(int)
	if !ok {
		return 0, fmt.Errorf("Data is not integer.")
	}
	q.slice = q.slice[1:len(q.slice)]
	return item, err
}

func (q *Queue) DequeueGeneric() (item interface{}, err error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("No elements present in queue")
	}
	item = q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return item, err
}

func (q *Queue) IsEmpty() bool {
	if len(q.slice) == 0 {
		return true
	}
	return false
}
