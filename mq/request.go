package mq

type Request struct {
	method string
	uri    string
	body   string
	next   *Request
}
