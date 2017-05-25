package handlers

import (
    "github.com/streadway/amqp"
    "log"
)

type Handler interface {
    // handler handles different type and its data from mq.
    // New must return a non-nil error if it could not been initialed
    Connect()(err error)
    // Listen the new message in mq
    Listen()(err error)
    // close the running handler
    CLose()(err error)
}

type Handler struct {
}

func (handler *Handler) Connect() (conn *amqp.Connection, err error) {
    if conn, err = amqp.Dial(address + ":" + port); err != nil {
        log.Fatalln("new error")
        return nil, err
    }
    return conn, err
}
