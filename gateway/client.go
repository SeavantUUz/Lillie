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
)

const (
    RUN = 0x0001
    CLOSE = 0x0002
)

type Client struct {
    uuid string
    in chan string
    out chan string
    reader *bufio.Reader
    writer *bufio.Writer
    status int
    quit chan bool
}

func NewClient(conn *net.Conn) *Client {
    uuid := tool.UUID()
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)
    client := &Client{
        uuid: uuid,
        in: make(chan protocol.Request),
        out: make(chan protocol.Response),
        quit: make(chan bool),
        reader: reader,
        writer: writer,
        status:int{},
    }
    client.Listen()
    return client
}

func (client *Client) Listen() {
    client.status = RUN
    defer client.Close()
    go func() {
        for {
            select {
            case <- client.quit:
                return
            case response := <- client.out:
                client.Write(&response)
            case request := <- client.in :
                client.Dispatch(&request)
            }
        }
    }()
}

func (client *Client) Close() {
    client.status = CLOSE
    log.Println("client exit")
    client.quit <- true
}

func (client *Client) Read() (err error) {
    msgLen, err := binary.ReadUvarint(client.reader)
    if err != nil {
        log.Println("Read Error")
        client.quit <- true
        return err
    }
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
    protoSize := uint64(len(out))
    bytes := make([]byte, len(protoSize) + protoSize)
    n := binary.PutUvarint(bytes, protoSize)
    copy(bytes[n:], out)
    client.writer.Write(bytes)
    client.writer.Flush()
    return nil
}

func (client *Client) Dispatch(request *protocol.Request) (err error) {
    operation := request.Operation
    switch operation {
    case protocol.Operation_MESSAGE_SEND:
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
    }
    client.out <- response
    return nil
}
