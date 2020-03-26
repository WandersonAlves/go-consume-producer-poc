package main

import (
	"go-consumer-producer-poc/amqp"
	"go-consumer-producer-poc/shared"
)

const queueName = "queue-test"

func main() {
	shared.Print("Go Consumer")
	shared.Print("Trying to connect to RabbitMQ...")
	connection, err := amqp.ConnectToAMQP()
	shared.FatalError(err, "ConnectionException")
	shared.Print("Trying to assert queue...")
	channel, routingKey, err := amqp.ConnectToChannelAndAssertQueue(connection, queueName)
	shared.FatalError(err, "ChannelAndQueueException")
	amqp.ConsumeMessages(channel, routingKey)
}
