package handlers

import (
    "github.com/streadway/amqp"
    "log"
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/golang/protobuf/proto"
    "github.com/garyburd/redigo/redis"
    "github.com/SeavantUUz/Lillie/tool"
    "github.com/SeavantUUz/Lillie/connector"
    "github.com/SeavantUUz/Lillie/configs"
)

type StoreHandler struct {
    base *Handler
    msgs chan *amqp.Delivery
    redis_pool *redis.Pool
}

func (handler *StoreHandler) Listen() (err error) {
    conn, ch, msgs, err := connector.PrepareMqConsume(configs.UPROUTER, "handler:store",
	    tool.RequestKey(protocol.Operation_MESSAGE_SEND))
    if err != nil {
        log.Fatalln("fail to launch a consumer")
    }
    defer conn.Close()
    defer ch.Close()
    handler.msgs = msgs
    
    forever := make(chan bool)
    go func() {
        for msg := range handler.msgs{
            go handler.store(msg)
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
    conn := handler.redis_pool.Get()
    defer conn.Close()
    conn.Do("HSET", "entity", tool.InboxKey(targetId), request)
    conn.Do("ZADD", "inbox", request.MsgId, tool.InboxKey(targetId)) // zadd field score content
    return nil
}

