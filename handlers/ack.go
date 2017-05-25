package handlers

import (
    "github.com/streadway/amqp"
    "log"
    "strconv"
    "github.com/SeavantUUz/Lillie/protocol"
)

type AckHandler struct {
    base *Handler
    msgs chan *amqp.Delivery
}

func (ack_handler *AckHandler) Listen() (err error) {
    defer ack_handler.Close()
    var conn *amqp.Connection
    if conn, err == ack_handler.base.Connect(); err != nil {
        log.Fatalln("fail to connect")
        return err
    }
    defer conn.Close()
    ch, err := conn.Channel()
    defer ch.Close()
    if err != nil {
        log.Fatalln("fail to open a channel")
        return err
    }
    q, err := ch.QueueDeclare(
        "receiver:"+strconv.FormatInt(int64(protocol.Operation_MESSAGE_ACK), 10),
        false,
        false,
        false,
        false,
        nil,
    )
    
    if err != nil {
        log.Fatalln("fail to declare a queue")
        return err
    }
    
    ack_handler.msgs, err == ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    
    if err != nil {
        log.Fatalln("fail to launch a consumer")
    }
    go func() {
        for msg := range ack_handler.msgs{
            ack_handler.reply(msg)
        }
    }()
    return nil
}

func (ack_handler *AckHandler) Close()  {
    close(ack_handler.msgs)
    log.Println("close ack handler")
}