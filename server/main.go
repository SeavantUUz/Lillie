package main
 //A router is important, but I delayed.

import (
    "net"
    "log"
    "github.com/SeavantUUz/Lillie/gateway"
    "fmt"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "7453"
    CONN_TYPE = "tcp"
)

type Server struct {
    clients map[string]*gateway.Client
    new_connection chan net.Conn
    exit chan bool
}

func NewServer() *Server {
    server:= &Server{
        clients:make(map[string]*gateway.Client),
        new_connection: make(chan net.Conn),
        exit: make(chan bool),
    }
    
    server.Run()
    
    return server
}

func (server *Server) Run() {
    go func() {
        for {
            select {
            case conn := <- server.new_connection:
                fmt.Println("new connection added")
                server.Add(conn)
            case <- server.exit:
                log.Println("gateway exited")
                return
            }
        }
    }()
}

func (server *Server) Add(conn net.Conn) *gateway.Client {
    client := gateway.NewClient(conn)
    uuid := client.GetUuid()
    server.clients[uuid] = client
    return client
}


func main() {
    server := NewServer()
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        server.exit <- true
        log.Printf("error %s occured when is listening", err)
    }
    for {
        conn, _ := l.Accept()
        server.new_connection <- conn
    }
    
}