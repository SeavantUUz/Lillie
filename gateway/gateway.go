package gateway

import (
    "net"
    "fmt"
)

type Gateway struct {
    receivers map[string]*Receiver
    sender *Sender
    exit chan bool
    leave chan string
}

func NewGateway() *Gateway{
    server:= &Gateway{}
    sender := NewSender(server)
    server.sender = sender
    server.receivers = make(map[string]*Receiver)
    server.exit = make(chan bool)
    server.leave = make(chan string)
    return server
}

func (server *Gateway) Run() {
    
    for {
        select {
        case <- server.exit:
            server.Stop()
            return
        case uuid := <- server.leave:
            server.Clean(uuid)
        }
    }
}

func (server *Gateway) Stop() {
    ch := make(chan bool)
    for _, receiver := range server.receivers {
        go func(receiver *Receiver) {
            receiver.quit <- true
        }(receiver)
    }
    for range server.receivers {
        <-ch
    }
    close(server.leave)
    close(server.exit)
}

func (server *Gateway) Clean(uuid string) {
    delete(server.receivers, uuid)
    fmt.Printf("receiver %s left\n", uuid)
}

func (server *Gateway) Join(conn net.Conn) {
    go func() {
        receiver := NewReceiver(conn, server)
        uuid := receiver.uuid
        server.receivers[uuid] = receiver
        receiver.run()
    }()
}

