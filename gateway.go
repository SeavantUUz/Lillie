package main
 //A router is important, but I delayed.

import (
    "net"
    "log"
    "github.com/SeavantUUz/Lillie/gateway"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "7453"
    CONN_TYPE = "tcp"
)


type Gateway struct {
    clients [] *gateway.Client
    new_connection chan net.Conn
    exit chan bool
}


func (gateway *Gateway) Run() {
    go func() {
        for {
            select {
            case conn := gateway.new_connection:
                gateway.Add(conn)
            case <- gateway.exit:
                log.Println("gateway exited")
                return
            }
        }
    }()
}

func (Gateway *Gateway) Add(conn *net.Conn) {

}

func NewGateWay() *Gateway {
    gateway := &Gateway{
        clients: make([] *gateway.Client, 0),
        new_connection: make(chan net.Conn),
        exit: make(chan bool),
    }
    
    gateway.Run()
    
    return gateway
}


func main() {
    gateway := NewGateWay()
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        log.Printf("error %s occured when is listening", err)
    }
    for {
        conn, err := l.Accept()
        in
    }
    
    
}