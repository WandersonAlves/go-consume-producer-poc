package main

import (
	"encoding/json"
	"fmt"
	"go-consumer-producer-poc/amqp"
	"go-consumer-producer-poc/structs"
	"log"
	"time"
)

const queueName = "queue-test"

func generateMessage(str string) []byte {
	message := structs.QueueMessage{
		Title:     str,
		Timestamp: time.Now().Unix(),
	}
	body, err := json.Marshal(message)
	fatalError(err, "JSONMarshlingException")
	return body
}

func print(str string) {
	fmt.Printf(str + "\n")
}

func printErrorIfExists(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func fatalError(err error, msg string) {
	if err != nil {
		log.Fatalf("[%s]: %s", msg, err)
	}
}

func main() {
	print("Go Producer")
	print("Trying to connect to RabbitMQ...")
	connection, err := amqp.ConnectToAMQP()
	fatalError(err, "ConnectionException")
	print("Trying to assert queue...")
	channel, routingKey, err := amqp.ConnectToChannelAndAssertQueue(connection, queueName)
	fatalError(err, "ChannelAndQueueException")
	message := generateMessage("Oisss")
	print("Sending message...")
	err = amqp.PublishMessage(channel, routingKey, message)
	printErrorIfExists(err)
}
