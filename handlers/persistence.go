package handlers

import (
    "github.com/streadway/amqp"
    "log"
    "strconv"
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/golang/protobuf/proto"
)

type PersistenceHandler struct {
    base *Handler
    msgs chan *amqp.Delivery
}

func (handler *PersistenceHandler) Listen() (err error) {
    var exchange_name string
    var queue_name string
    
    // what up event should be deal
    exchange_name = "up:" + strconv.FormatInt(int64(protocol.Operation_MESSAGE_SEND), 10)
    // the handler
    queue_name = "handler:persistence"
    
    defer handler.Close()
    var conn *amqp.Connection
    if conn, err == handler.base.Connect(); err != nil {
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
    
    err = ch.ExchangeDeclare(
        exchange_name,
        "fanout",
        true, // durable
        false, //auto_delete
        false, // internal
        false, // no wait
        nil,
    )
    
    if err != nil {
        log.Fatalln("fail to declare a exchange")
        return err
    }
    
    
    q, err := ch.QueueDeclare(
        queue_name,
        false, // durable
        false, // delete when unuse
        true, // exclusive
        false, // no-wait
        nil, // argument
    )
    
    if err != nil {
        log.Fatalln("fail to declare a queue")
        return err
    }
    
    err = ch.QueueBind(
        q.Name,
        "",
        exchange_name,
        false,
        nil,
    )
    
    if err != nil {
        log.Fatalln("fail to bind queue to exchange")
    }
    
    handler.msgs, err == ch.Consume(
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
    
    forever := make(chan bool)
    go func() {
        for msg := range handler.msgs{
            handler.reply(msg)
        }
    }()
    log.Println("[*] Waiting for logs. To exit press CTRL+C")
    <- forever
    return nil
}

func (handler *PersistenceHandler) Close()  {
    close(handler.msgs)
    log.Println("close persistence handler")
}

func (handler *PersistenceHandler) reply(msg *amqp.Delivery) (error) {
    request := &protocol.Request{}
    if err := proto.Unmarshal(msg.Body, request); err != nil {
        log.Fatalln("Failed to parse Request:", err)
        return err
    }
    log.Printf("Persistence handler received a request: %s", request)
    persistence_request(request)
    return nil
}

