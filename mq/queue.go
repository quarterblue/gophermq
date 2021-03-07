package mq

type Queue struct {
	head *Request
	tail *Request
	size int
}
