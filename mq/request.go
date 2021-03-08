package mq

type Request struct {
	method string
	uri    string
	body   string
	next   *Request
	prev   *Request
}

func createRequest(method string, uri string, body string) *Request {
	newRequest := &Request{
		method: method,
		uri:    uri,
		body:   body,
		next:   nil,
		prev:   nil,
	}

	return newRequest
}

func (req *Request) deleteRequest() {
}

func (req *Request) writeRequest() {
}
