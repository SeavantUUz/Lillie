package handlers

import (
    "github.com/streadway/amqp"
    "log"
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/golang/protobuf/proto"
    "strconv"
    "github.com/go-redis/redis"
)

type StoreHandler struct {
    base *Handler
    msgs chan *amqp.Delivery
    redis *redis.Conn
}

func (handler *StoreHandler) Listen() (err error) {
    queue_name := "handler:store"
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
        UPROUTER,
        "direct",
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
        "request:"+ strconv.FormatInt(int64(protocol.Operation_MESSAGE_SEND), 10),
        UPROUTER,
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
            handler.store(msg)
        }
    }()
    log.Println("[*] Waiting for logs. To exit press CTRL+C")
    <- forever
    return nil
}

func (handler *StoreHandler) Close()  {
    close(handler.msgs)
    log.Println("close store handler")
}

func (handler *StoreHandler) store(msg *amqp.Delivery) (error) {
    request := &protocol.Request{}
    if err := proto.Unmarshal(msg.Body, request); err != nil {
        log.Fatalln("Failed to parse Request:", err)
        return err
    }
    log.Printf("Receive a request: %s", request)
    targetId := request.TargetId
    return nil
}

