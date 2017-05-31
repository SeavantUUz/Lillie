package handlers

import (
    "github.com/streadway/amqp"
    "log"
    "github.com/garyburd/redigo/redis"
    "time"
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

type Handler struct {}

// Do you think these function should move from there to other certain directory?
func (handler *Handler) connect_to_mq() (conn *amqp.Connection, err error) {
    if conn, err = amqp.Dial(ADDRESS + ":" + PORT); err != nil {
        log.Fatalln("connect error", err)
        return nil, err
    }
    return conn, err
}


func (handler *Handler) connect_to_redis()  *redis.Pool  {
   return &redis.Pool{
       MaxIdle: 10,
       IdleTimeout: 240 * time.Second,
       Dial: func() (redis.Conn, error) {return redis.Dial("tcp", REDISADDR)},
   }
}

