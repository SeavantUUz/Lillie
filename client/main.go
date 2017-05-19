package main

import (
    "net"
    "log"
    "github.com/golang/protobuf/proto"
    "encoding/binary"
    "github.com/SeavantUUz/Lillie/protocol"
    "time"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "7453"
    CONN_TYPE = "tcp"
)

func varintSize(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func main(){
    conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        log.Printf("error %s occured when is listening", err)
    }
    for {
        content := &protocol.Send{
                Text: "你好",
        }
        body, err := proto.Marshal(content)
        if err != nil {
            log.Fatalln("Marshal error")
        }
        request := &protocol.Request{
            SourceId: 11111,
            TargetId: 22222,
            Operation:protocol.Operation_MESSAGE_SEND,
            MsgId: 123123123,
            Timestamp: uint64(time.Now().Unix()),
            Body: body,
        }
        out, err := proto.Marshal(request)
        if err != nil {
            log.Fatalln("Failed to encode reponse data", err)
        }
        protoSize := len(out)
        bytes := make([]byte, uint64(varintSize(uint64(protoSize)) + protoSize))
        n := binary.PutUvarint(bytes, uint64(protoSize))
        copy(bytes[n:], out)
        conn.Write(bytes)
        time.Sleep(1)
    }
}
