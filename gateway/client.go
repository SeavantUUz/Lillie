package gateway

import (
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/SeavantUUz/Lillie/tool"
    "github.com/golang/protobuf/proto"
    "log"
    "encoding/binary"
    "bufio"
    "net"
    "time"
    "fmt"
)


type Client struct {
    uuid string
    in chan *protocol.Request
    out chan *protocol.Response
    reader *bufio.Reader
    writer *bufio.Writer
    status int
    quit chan bool
    server *Server
}

func NewClient(conn net.Conn, server *Server) *Client {
    uuid := tool.UUID()
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)
    client := &Client{
        uuid: uuid,
        in: make(chan *protocol.Request),
        out: make(chan *protocol.Response),
        quit: make(chan bool),
        reader: reader,
        writer: writer,
        status: INIT,
        server: server,
    }
    return client
}


func (client *Client) GetUuid() string {
    return client.uuid
}

// it's not good way to mix fuck'in jobs in one function
func (client *Client) Run() {
    client.status = RUN
    client.Listen()
}

func (client *Client) Listen()  {
    go func() {
        for client.status == RUN {
            client.Read()
        }
    }()
    
    go func() {
        for response :=  range client.out {
            client.Write(response)
        }
    }()
    
    go func() {
        for request := range client.in {
            client.Dispatch(request)
        }
    }()
    
    go func() {
        for _ = range client.quit {
            fmt.Println("quit")
            client.Close()
        }
    }()
}

func (client *Client) Close() {
    client.status = CLOSE
    close(client.in)
    close(client.out)
    close(client.quit)
    client.server.leave <- client.uuid
}

func (client *Client) Read() (err error) {
    msgLen, err := binary.ReadUvarint(client.reader)
    if err != nil {
        log.Println("Read Error")
        client.Close()
        return err
    }
    fmt.Println(msgLen)
    bytes := make([]byte, msgLen)
    client.reader.Read(bytes)
    request := &protocol.Request{}
    if err := proto.Unmarshal(bytes, request); err != nil {
        log.Fatalln("Failed to parse request data:", err)
        return err
    }
    client.in <- request
    return nil
}


func (client *Client) Write(response *protocol.Response) (err error) {
    out, err := proto.Marshal(response)
    if err != nil {
        log.Fatalln("Failed to encode reponse data", err)
        return err
    }
    protoSize := len(out)
    bytes := make([]byte, uint64(tool.VarIntSize(uint64(protoSize)) + protoSize))
    n := binary.PutUvarint(bytes, uint64(protoSize))
    copy(bytes[n:], out)
    client.writer.Write(bytes)
    client.writer.Flush()
    return nil
}

func (client *Client) Dispatch(request *protocol.Request) (err error) {
    operation := request.Operation
    switch operation {
    case protocol.Operation_MESSAGE_SEND:
        fmt.Println("send")
        if err := client.ack(request); err != nil {
            return err
        }
    }
    return nil
}


// private methods
func (client *Client) ack(request *protocol.Request) (err error)  {
    msgId := request.MsgId
    response := &protocol.Response{
        MsgId: msgId ,
        Timestamp: uint64(time.Now().Unix()),
        Operation: protocol.Operation_MESSAGE_ACK,
        Body: []byte{},
    }
    client.out <- response
    return nil
}
