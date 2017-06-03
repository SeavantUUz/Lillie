package main

import (
    "net"
    "log"
    "encoding/binary"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "7454"
    CONN_TYPE = "tcp"
)

func main() {
    var num int64
    num = 0
    listener, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if err != nil {
        log.Fatal("numnode starts up error", err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("accept error", err)
        }
        go handleConn(conn, num)
        num += 1
    }
}

func handleConn(conn net.Conn, num uint16) {
    defer conn.Close()
    binary.Write(conn, binary.LittleEndian, num)
}

