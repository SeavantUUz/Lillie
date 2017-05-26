package trials


import (
    "github.com/streadway/amqp"
    "log"
    "strconv"
    "github.com/SeavantUUz/Lillie/protocol"
)

func failOnError(err error, msg string){
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}


func main() {
    conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672")
    defer conn.Close()
    ch, err := conn.Channel()
    defer ch.Close()
    
    exchange_name := strconv.FormatInt(int64(protocol.Operation_MESSAGE_SEND), 10)
    queue_name := "handler:" + strconv.FormatInt(int64(protocol.Operation_MESSAGE_ACK), 10)
    
    failOnError(err, "Failed to open a channel")
    
    err := ch.ExchangeDeclare(
        exchange_name,
        "fanout",
        true, // durable
        false, //auto_delete
        false, // internal
        false, // no wait
        nil,
    )
    
    q, err := ch.QueueDeclare(
        queue_name,
        false, // durable
        false, // delete when unuse
        true, // exclusive
        false, // no-wait
        nil, // argument
    )
    
    failOnError(err, "Failed to declare a queue")
    
    body := "hello"
    err = ch.Publish(
        "",
        q.Name,
        false,
        false,
        amqp.Publishing{
            ContentType:"text/plain",
            Body: []byte(body),
        })
    failOnError(err, "Failed to publish a message")
    log.Printf(" [x] Sent %s", body)
}