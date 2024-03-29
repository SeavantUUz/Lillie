package handlers

import (
    "github.com/streadway/amqp"
    "log"
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/golang/protobuf/proto"
    "github.com/SeavantUUz/Lillie/tool"
	"database/sql"
	_ "github.com/lib/pq"
    "github.com/SeavantUUz/Lillie/connector"
    "github.com/SeavantUUz/Lillie/configs"
)

type PersistenceHandler struct {
    base *Handler
    msgs chan *amqp.Delivery
}

func (handler *PersistenceHandler) Listen() (err error) {
    conn, ch, msgs, err := connector.PrepareMqConsume(configs.UPROUTER, "handler:persistence",
	    tool.RequestKey(protocol.Operation_MESSAGE_SEND))
    if err != nil {
        log.Fatalln("fail to launch a consumer")
    }
    defer conn.Close()
    defer ch.Close()
    handler.msgs = msgs
    
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
        return err
    }
    
    forever := make(chan bool)
    go func() {
        for msg := range handler.msgs{
            go handler.store(msg) // I have put the method into Handler, but I thought it was not a good way.
        }
    }()
    log.Println("[*] Waiting for logs. To exit press CTRL+C")
    <- forever
    return nil
}

func (handler *PersistenceHandler) Close()  {
    close(handler.msgs)
    log.Println("close store handler")
}

func (handler *PersistenceHandler) store(msg *amqp.Delivery) (err error) {
    request := &protocol.Request{}
    if err = proto.Unmarshal(msg.Body, request); err != nil {
        log.Fatalln("Failed to parse Request:", err)
        return err
    }
    log.Printf("Receive a request: %s", request)
    go handler.persistence(request)
    if err != nil {
        log.Fatalln("Persistence request failed:", request)
        return err
    }
    return nil
}

func (handler *PersistenceHandler) persistence(request *protocol.Request) (err error) {
    db, err := sql.Open("postgres", PGCONFIG)
    if err != nil {
        log.Fatalln("connect to pg failed")
        return err
    }
    stmt, err := db.Prepare("insert into message (msgid, seq, sourceid, targetid, operation, content) VALUES " +
            "($1, $2, $3, $4, $5, $6)")
    if err != nil {
        log.Fatalln("failed to get a sql connection")
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(request.MsgId, request.Seq, request.SourceId, request.TargetId, request.Operation,
        request.Body)
    if err != nil {
        log.Fatalln("Failed to insert into a message", request)
        return err
    }
    return nil
}

