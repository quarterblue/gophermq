package mq

type MessageQueue struct {
	name string
	host string
	port string

	outgoing *Queue
	incoming *Queue
	shutdown bool
}

func createMq(name string, host string, port string) *MessageQueue {
	newMQ := &MessageQueue{
		name:     name,
		host:     host,
		port:     port,
		outgoing: &Queue{},
		incoming: &Queue{},
		shutdown: false,
	}

	return newMQ
}
