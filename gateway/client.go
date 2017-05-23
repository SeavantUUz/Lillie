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

const (
    INIT = 0x0000
    RUN = 0x0001
    CLOSE = 0x0002
)

type Client struct {
    uuid string
    in chan *protocol.Request
    out chan *protocol.Response
    reader *bufio.Reader
    writer *bufio.Writer
    status int
    quit chan bool
}

func NewClient(conn net.Conn) *Client {
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
    }
    return client
}


func (client *Client) GetUuid() string {
    return client.uuid
}

func (client *Client) Run() {
    client.status = RUN
    go client.poll()
    for client.status == RUN {
        client.Read()
    }

}

func (client *Client) poll()  {
    for {
        select {
        case <- client.quit:
            client.Close()
            return
        case response := <- client.out:
            fmt.Println("shit")
            client.Write(response)
        case request := <- client.in :
            client.Dispatch(request)
            //fmt.Println(request)
        }
    }
}

func (client *Client) Close() {
    client.status = CLOSE
    log.Println("client exit")
}

func (client *Client) Read() (err error) {
    msgLen, err := binary.ReadUvarint(client.reader)
    if err != nil {
        log.Println("Read Error")
        client.quit <- true
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
    fmt.Print(request)
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
    fmt.Println("out2")
    client.out <- response
    fmt.Println("out")
    return nil
}
