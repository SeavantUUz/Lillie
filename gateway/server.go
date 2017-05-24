package gateway

import (
    "net"
    "fmt"
)



type Server struct {
    clients map[string]*Client
    new_connection chan net.Conn
    exit chan bool
    leave chan string
    status int
}

func NewServer() *Server {
    server:= &Server{
        clients:make(map[string]*Client),
        exit: make(chan bool),
        leave: make(chan string),
        status: INIT,
    }
    
    return server
}

func (server *Server) Run() {
    server.status = RUN
    for server.status == RUN {
        select {
        case <- server.exit:
            server.Stop()
            return
        case uuid := <- server.leave:
            server.Clean(uuid)
        }
    }
}

func (server *Server) Stop() {
    server.status = CLOSE
    close(server.new_connection)
    ch := make(chan bool)
    for _, v := range server.clients {
        go func(v *Client) {
            v.status = CLOSE
            v.quit <- true
        }(v)
    }
    for range server.clients {
        <-ch
    }
    close(server.leave)
    close(server.exit)
}

func (server *Server) Clean(uuid string) {
    delete(server.clients, uuid)
    fmt.Printf("client %s left\n", uuid)
}

func (server *Server) Join(conn net.Conn) {
    go func() {
        client := NewClient(conn, server)
        uuid := client.GetUuid()
        server.clients[uuid] = client
        client.Run()
    }()
}

