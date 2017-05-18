package trials


import (
    "github.com/streadway/amqp"
    "log"
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
    
    failOnError(err, "Failed to open a channel")
    
    q, err := ch.QueueDeclare(
        "hello",
        false,
        false,
        false,
        false,
        nil,
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