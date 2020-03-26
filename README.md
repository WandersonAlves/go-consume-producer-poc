# go-consumer-producer-poc

This PoC aims to create a RabbitMQ producer and consumer

The type of the messages to be produced are:

```go
// QueueMessage defines the message to be sent to server
type QueueMessage struct {
	Title     string
	Timestamp int32
}
```

# How to run

First, start your RabbitMQ instance locally using docker running `docker run -d -p 8080:15672 -p 5672:5672 -p 25676:25676 rabbitmq:3-management`

Run `go run ./goconsumer/main.go` to start the consumer. The consumer will keep listen to events on `test-queue` that will be produced by `goproducer`.

Then, run `go run ./goproducer/main.go`. This'll produce a message in the queue and exit.

Meanwhile, check your RabbitMQ UI acessing http://localhost:8080/#/queues/%2F/queue-test