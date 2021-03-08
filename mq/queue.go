package mq

import (
	"errors"
	"sync"
)

type Queue struct {
	sync.Mutex
	head *Request
	tail *Request
	size uint32
}

// Constructor function for the Queue data struct
func createQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Enqueue: Push request to the back of the queue
func (q *Queue) pushQueue(r *Request) error {
	q.Lock()

	if q.size == 0 {
		q.head = r
		q.tail = r
		q.Unlock()
		return nil
	}

	tailQueue := q.tail
	tailQueue.next = r
	q.tail = r
	q.Unlock()
	return nil
}

// Dequeue: request from the front of the queue
func (q *Queue) popQueue() (*Request, error) {
	q.Lock()
	if q.size == 0 {
		err := errors.New("Queue is empty")
		q.Unlock()
		return nil, err
	}

	// TODO: logic for popping the queue
	var poppedRequest *Request
	if q.size == 1 {
		// Assert that head is equal to tail to verify the integrity of the queue
		if q.head != q.tail {
			err := errors.New("Data structure invalid")
			return nil, err
		}
		// Remove the item and set head and tail to nil
		poppedRequest = q.head
		q.head = nil
		q.tail = nil
		q.size--
	}

	// General case logic
	poppedRequest = q.head
	q.head = poppedRequest.next
	poppedRequest.next = nil
	poppedRequest.prev = nil
	q.Unlock()

	return poppedRequest, nil
}
