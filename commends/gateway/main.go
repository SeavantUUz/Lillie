package main

import (
    "net"
    "log"
    "github.com/SeavantUUz/Lillie/gateway"
    "os"
    "os/signal"
    "syscall"
    "fmt"
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
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
    go func() {
        <-c
        go server.Stop()
        os.Exit(1)
        fmt.Println("hahaha")
    }()
    for {
        conn, _ := l.Accept()
        server.Join(conn)
    }
}
