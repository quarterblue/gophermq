package mq

import "errors"

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
		outgoing: nil,
		incoming: nil,
		shutdown: false,
	}

	return newMQ
}

func deleteMq(mq *MessageQueue) {

}

func (mq *MessageQueue) publishMq(topic string, body string) error {
	return errors.New("Place Holder Error")
}

func (mq *MessageQueue) retrieveMq() error {
	return errors.New("Place Holder Error")
}

func (mq *MessageQueue) subscribeMq(topic string) error {
	return errors.New("Place Holder Error")
}

func (mq *MessageQueue) unsubscribeMq(topic string) error {
	return errors.New("Place Holder Error")
}

func (mq *MessageQueue) startMq() {

}

func (mq *MessageQueue) stopMq() {

}

func (mq *MessageQueue) shutdownMq() bool {
	return true
}
