package handlers

import (
    "github.com/streadway/amqp"
    "log"
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/golang/protobuf/proto"
    "github.com/SeavantUUz/Lillie/connector"
    "github.com/SeavantUUz/Lillie/tool"
    "github.com/SeavantUUz/Lillie/configs"
)

type AckHandler struct {
    base *Handler
    msgs chan *amqp.Delivery
}

func (handler *AckHandler) Listen() (err error) {
    conn, ch, msgs, err := connector.PrepareMqConsume(configs.UPROUTER, "handler:ack",
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
            handler.reply(msg)
        }
    }()
    log.Println("[*] Waiting for logs. To exit press CTRL+C")
    <- forever
    return nil
}

func (handler *AckHandler) Close()  {
    close(handler.msgs)
    log.Println("close ack handler")
}

func (handler *AckHandler) reply(msg *amqp.Delivery) (error) {
    request := &protocol.Request{}
    if err := proto.Unmarshal(msg.Body, request); err != nil {
        log.Fatalln("Failed to parse Request:", err)
        return err
    }
    log.Printf("Receive a request: %s", request)
    seq := request.Seq
    msgId := request.MsgId
    sourceId := request.SourceId
    response := &protocol.Response{
        SourceId: 0, // the sender of ack is server, so assigned sourceId to 0
        TargetId:sourceId,
        MsgId: msgId,
        Seq: seq,
        Router: request.Router,
        Operation: protocol.Operation_MESSAGE_ACK,
    }
    connector.Down(response)
    return nil
}

