package mq

import "errors"

type Queue struct {
	head *Request
	tail *Request
	size uint32
}

func createQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue) pushQueue(r *Request) error {
	if q.size == 0 {
		q.head = r
		q.tail = r
		return nil
	}

	tailQueue := q.tail
	tailQueue.next = r
	q.tail = r
	return nil
}

func (q *Queue) popQueue() (*Request, error) {
	if q.size == 0 {
		err := errors.New("Queue is empty")
		return nil, err
	}

	// TODO: logic for popping the queue
	return nil, nil
}
