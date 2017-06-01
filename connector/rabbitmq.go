package connector

import (
    "github.com/SeavantUUz/Lillie/protocol"
    "github.com/streadway/amqp"
    "log"
    "github.com/golang/protobuf/proto"
    "github.com/SeavantUUz/Lillie/tool"
    "github.com/SeavantUUz/Lillie/handlers"
    "github.com/SeavantUUz/Lillie/configs"
)

func PrepareMqConsume(exchange_name, queue_name, routing_key string) (*amqp.Connection, *amqp.Channel, <-chan amqp.Delivery, error) {
    var mq_config configs.MQConfig
    conn, err := amqp.Dial(mq_config.Address + ":" + mq_config.Port)
    if err != nil {
        log.Fatalln("Dial Mq Failed", err)
        return nil, nil, nil, err
    }
    ch, err := conn.Channel()
    if err != nil {
        log.Fatalln("fail to open a channel")
        return conn, nil, nil, err
    }
    
    err = ch.ExchangeDeclare(
        exchange_name,
        "direct",
        true, // durable
        false, //auto_delete
        false, // internal
        false, // no wait
        nil,
    )
    if err != nil {
        log.Fatalln("fail to declare a exchange")
        return conn, ch, nil, err
    }
    
    
    q, err := ch.QueueDeclare(
        queue_name,
        false, // durable
        false, // delete when unuse
        true, // exclusive
        false, // no-wait
        nil, // argument
    )
    
    if err != nil {
        log.Fatalln("fail to declare a queue")
        return conn, ch, nil, err
    }
    
    err = ch.QueueBind(
        queue_name,
        //tool.RequestKey(protocol.Operation_MESSAGE_SEND),
        routing_key,
        exchange_name,
        false,
        nil,
    )
    
    if err != nil {
        log.Fatalln("fail to bind queue to exchange")
        return conn, ch, nil, err
    }
    
    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    
    return conn, ch, msgs, nil
}

// the two function below are the same actually.
// up leads request from gateway to handler
// down returns response back to gateway

// I knew the two functions could merge to one. But I didn't want to do that.
// there is a topic should be discuss, shall the down function need a priority declare?
// I think not

// throw request to exchange handlers_uprouter
// message hang out to different queue by routing key
func Up(request *protocol.Request) {
    conn, err := amqp.Dial("amqp://guest:guest@172.17.0.3:5672/")
    if err != nil {
        log.Fatalln("Failed to connect to RabbitMQ", err)
        return
    }
    defer conn.Close()
    ch, err := conn.Channel()
    defer ch.Close()
    exchange_name := handlers.UPROUTER
    err = ch.ExchangeDeclare(
        exchange_name,
        "direct",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalln("Failed to declare exchange")
        return
    }
    out, err := proto.Marshal(request)
    if err != nil {
        log.Fatalln("Failed to marsh a request", request)
        return
    }
    err = ch.Publish(
        exchange_name, // exchange_name
        tool.RequestKey(request.Operation), //routing key
        false,
        false,
        amqp.Publishing{
            ContentType:"text/plain",
            Body:[]byte(out),
        })
    if err != nil {
        log.Fatalln("Failed to publish message", request)
        return
    }
    log.Println("up request sent", request)
}


// the queue should be subscribe by one of gateway's processes
func Down(response *protocol.Response) {
    operation := response.Operation
    conn, err := amqp.Dial("amqp://guest:guest@172.17.0.3:5672/")
    if err != nil {
        log.Fatalln("Failed to connect to RabbitMQ", err)
        return
    }
    defer conn.Close()
    ch, err := conn.Channel()
    var priority uint8
    switch operation {
    case protocol.Operation_MESSAGE_ACK:
        priority = 0
    case protocol.Operation_MESSAGE_NOTIFY:
        priority = 9
    default:
        priority = 0
    }
    defer ch.Close()
    exchange_name := handlers.DOWNROUTER
    var args amqp.Table
    args["x-max-priority"] = 10
    err = ch.ExchangeDeclare(
        exchange_name,
        "direct",
        true,
        false,
        false,
        false,
        args,
    )
    if err != nil {
        log.Fatalln("Failed to declare exchange", err)
        return
    }
    out, err := proto.Marshal(response)
    if err != nil {
        log.Fatalln("Failed to marsh a response", err)
        return
    }
    
    err = ch.Publish(
        "",
        exchange_name,
        false,
        false,
        amqp.Publishing{
            ContentType:"text/plain",
            Body:[]byte(out),
            Priority: priority,
        })
    if err != nil {
        log.Fatalln("Failed to publish message", response)
        return
    }
    log.Println("down response", response)
}
