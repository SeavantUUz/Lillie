package gateway

import (
    "github.com/SeavantUUz/Lillie/handlers"
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/golang/protobuf/proto"
    "log"
    "github.com/SeavantUUz/Lillie/connector"
)

type Sender struct {
    quit chan bool
    server *Gateway
}

func NewSender(server *Gateway) *Sender {
    quit := make(chan bool)
    s := &Sender{
        quit:quit,
        server: server,
    }
    return s
}

func (s *Sender) run() error {
    conn, ch, msgs, err := connector.PrepareMqConsume(handlers.DOWNROUTER, "", "") // consume all message in downrouter exchange
    defer s.stop()
    defer conn.Close()
    defer ch.Close()
    if err != nil {
        log.Fatalln("Failed to prepare mq", err)
        return err
    }
    go func() {
        for msg := range msgs{
            response := &protocol.Response{}
            if err = proto.Unmarshal(msg.Body, response); err != nil {
                log.Fatalln("Failed to parse Response:", err)
                return err
            }
            targetid := response.TargetId
            router := &response.Router
            s.send(targetid, router, response)
        }
    }()
    <- s.quit
    return nil
}

func (s *Sender) stop() error {
    return nil
}

func (s *Sender) send(targetid uint64, router *protocol.Router, response *protocol.Response) {
    return nil
}


