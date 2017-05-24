package main

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

func main() {
    server := gateway.NewServer()
    go server.Run()
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        go server.Stop()
        log.Printf("error %s occured when is listening", err)
    }
    for {
        conn, _ := l.Accept()
        server.Join(conn)
    }
    server.Stop()
}
