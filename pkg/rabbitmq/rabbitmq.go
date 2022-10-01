package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// Opens and returns a communication channel with Rabbitmq
func OpenChannel() (*amqp.Channel, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	rabbitmqChannel, err := connection.Channel()
	rabbitmqChannel.Qos(100, 0, false)

	if err != nil {
		panic(err)
	}

	return rabbitmqChannel, nil
}

// Consumes every message that arrives at Rabbitmq in the order queue, and writes that message to a channel
func Consume(rabbitmqChannel *amqp.Channel, messageDeliveryChannel chan<- amqp.Delivery) error {
	messages, err := rabbitmqChannel.Consume(
		"orders",      // queue
		"go-consumer", // consumer
		false,         // autoAck
		false,         // exclusive
		false,         // noLocal
		false,         // noWait
		nil,           // args
	)

	if err != nil {
		return err
	}

	for message := range messages {
		messageDeliveryChannel <- message
	}

	return nil
}
