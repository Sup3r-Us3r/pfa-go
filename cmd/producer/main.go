package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	Id    string
	Price float64
}

// Generates a random fake Order
func generateOrder() Order {
	return Order{
		Id:    uuid.New().String(),
		Price: rand.Float64() * 100.0,
	}
}

// Send a message to Rabbitmq amq.direct exchange with the order data
func Notify(rabbitmqChannel *amqp.Channel, order Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	body, err := json.Marshal(order)

	if err != nil {
		return err
	}

	err = rabbitmqChannel.PublishWithContext(
		ctx,          // context
		"amq.direct", // exchange
		"",           // key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: 2,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

// Produces 1.000 messages and sends them to Rabbitmq
func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	rabbitmqChannel, err := connection.Channel()

	if err != nil {
		panic(err)
	}

	defer rabbitmqChannel.Close()

	for i := 0; i < 1000; i++ {
		order := generateOrder()
		err := Notify(rabbitmqChannel, order)

		if err != nil {
			panic(err)
		}
	}
}
