package handlers

import (
    "github.com/streadway/amqp"
    "log"
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/golang/protobuf/proto"
    "strconv"
    "github.com/SeavantUUz/Lillie/connector"
    "github.com/SeavantUUz/Lillie/tool"
    "github.com/SeavantUUz/Lillie/configs"
)

type NotifyHandler struct {
    base *Handler
    msgs chan *amqp.Delivery
}

func (handler *NotifyHandler) Listen() (err error) {
    conn, ch, msgs, err := connector.PrepareMqConsume(configs.UPROUTER, "handler:notify",
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

func (handler *NotifyHandler) Close()  {
    close(handler.msgs)
    log.Println("close notify handler")
}

func (handler *NotifyHandler) reply(msg *amqp.Delivery) (error) {
    request := &protocol.Request{}
    if err := proto.Unmarshal(msg.Body, request); err != nil {
        log.Fatalln("Failed to parse Request:", err)
        return err
    }
    log.Printf("Receive a request: %s", request)
    targetId := request.TargetId
    lastMsgId := FetchLastMsgId(targetId) // fetch last msgId from user inbox
    notify := &protocol.Notify{
        LastMsgId: lastMsgId,
    }
    body, err := proto.Marshal(notify)
    if err != nil {
        log.Fatalln("protobuf marsh notify error", err)
        return err
    }
    response := &protocol.Response{
        SourceId: 0,
        TargetId:targetId,
        MsgId: 0,
        Seq: request.Seq,
        Router: request.Router,
        Operation: protocol.Operation_MESSAGE_NOTIFY,
        Body: body,
    }
    connector.Down(response)
    return nil
}

