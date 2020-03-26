package amqp

import (
	"fmt"
	"go-consumer-producer-poc/shared"

	"github.com/streadway/amqp"
)

const rabbitMQURL = "amqp://localhost:5672"

/*
ConsumeMessages receives a channel and a queueName and starts to consuming their contents
*/
func ConsumeMessages(channel *amqp.Channel, queueName string) error {
	messageChannel, err := channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	stopChan := make(chan bool)
	go func() {
		shared.Print("Ready to consume messages")
		for msg := range messageChannel {
			fmt.Printf("Received message with contents: %s\n", msg.Body)
			if err := msg.Ack(false); err != nil {
				fmt.Printf("Error acknowledging message : %s\n", err)
			}
		}
	}()
	<-stopChan

	return nil
}

/*
PublishMessage publishes a new message into a queue
*/
func PublishMessage(channel *amqp.Channel, queueName string, body []byte) error {
	message := amqp.Publishing{
		ContentType:  "text/plain",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	}
	if err := channel.Publish("", queueName, false, false, message); err != nil {
		return fmt.Errorf("Publish Error: %s", err)
	}

	return nil
}

/*
ConnectToChannelAndAssertQueue self explanatory
*/
func ConnectToChannelAndAssertQueue(connection *amqp.Connection, queueName string) (channel *amqp.Channel, routingKey string, err error) {

	channel, err = connection.Channel()

	if err != nil {
		defer connection.Close()
		return nil, "", fmt.Errorf("Can't create a channel: %s", err)
	}

	q, err := channel.QueueDeclare(queueName, true, false, false, false, nil)

	if err != nil {
		defer connection.Close()
		return nil, "", fmt.Errorf("Can't create a queue: %s", err)
	}

	return channel, q.Name, nil
}

/*
ConnectToAMQP self explanatory

Returns a new connect to a AMQP server
*/
func ConnectToAMQP() (*amqp.Connection, error) {
	connection, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("Can't connect to rabbitmq: %s", err)
	}

	return connection, nil
}
