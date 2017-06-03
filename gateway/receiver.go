package gateway

import (
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/SeavantUUz/Lillie/tool"
    "github.com/golang/protobuf/proto"
    "log"
    "net"
    "bytes"
    "io"
    "github.com/SeavantUUz/Lillie/connector"
)


type Receiver struct {
    uuid string
    in chan *protocol.Request
    out chan *protocol.Response
    conn net.Conn
    quit chan bool
    server *Gateway
}

func NewReceiver(conn net.Conn, server *Gateway) *Receiver {
    uuid := tool.UUID()
    client := &Receiver{
        uuid: uuid,
        in: make(chan *protocol.Request),
        out: make(chan *protocol.Response),
        quit: make(chan bool),
        conn: conn,
        server: server,
    }
    return client
}


func (r *Receiver) run()  {
    go r.read()
    for {
        select {
        case <- r.quit:
            r.stop()
            log.Println("quit client")
            return
        case response := <- r.out:
            go r.write(response)
        case request := <- r.in:
            go r.dispatch(request)
        }
    }
}

func (r *Receiver) stop() {
    r.conn.Close()
    r.server.leave <- r.uuid
}

func (r *Receiver) read() (err error) {
    var buf bytes.Buffer
    io.Copy(&buf, r.conn)
    request := &protocol.Request{}
    if err = proto.Unmarshal(buf, request); err != nil {
        log.Fatalln("Failed to parse request data:", err)
        return err
    }
    r.in <- request
    return nil
}


func (r *Receiver) write(response *protocol.Response) (err error) {
    out, err := proto.Marshal(response)
    if err != nil {
        log.Fatalln("Failed to encode reponse data", err)
        return err
    }
    _, err = io.Copy(r.conn, &out)
    if err != nil {
        log.Fatalln("Failed to copy data to conn", err)
        return err
    }
    return nil
}

func (r *Receiver) dispatch(request *protocol.Request)  {
    // add unique msg id
    msgId := r.server.node.Generate()
    request.Router = &protocol.Router{
        ClientId:r.uuid,
    }
    request.MsgId = msgId
    connector.Up(request)
}

