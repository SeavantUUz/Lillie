package tool

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

func GetNodeNum() int64 {
    conn, err := net.Dial(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if err != nil {
        log.Fatal(err)
        return binary.MaxVarintLen64
    }
    defer conn.Close()
    var num int64
    binary.Read(conn, binary.LittleEndian, &num)
    return num
}
