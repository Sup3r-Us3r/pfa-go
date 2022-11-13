package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/Sup3r-Us3r/pfa-go/internal/order/infra/database"
	"github.com/Sup3r-Us3r/pfa-go/internal/order/usecase"
	"github.com/Sup3r-Us3r/pfa-go/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Work to process the order that was read from the message channel
func worker(
	workerId int,
	messageDeliveryChannel <-chan amqp.Delivery,
	useCase *usecase.CalculateFinalPriceUseCase,
) {
	for message := range messageDeliveryChannel {
		var orderInput usecase.OrderInputDTO

		err := json.Unmarshal(message.Body, &orderInput)

		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}

		orderInput.Tax = 10.0

		_, err = useCase.Execute(orderInput)

		if err != nil {
			fmt.Println("Error executing use case", err)
		}

		message.Ack(false)

		println("WORK ID:", workerId, "ORDER PROCESSED:", orderInput.Id)
	}
}

// Start Database, HTTP Server, Rabbitmq Consumer and Work
func main() {
	const maxWorkers int = 3
	wg := sync.WaitGroup{}

	rootDirectory, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	db, err := sql.Open(
		"sqlite3",
		path.Join(
			rootDirectory, "internal", "order", "infra", "database", "sqlite.db",
		),
	)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)
	calculateFinalPriceUseCase := usecase.NewCalculateFinalPriceUseCase(repository)

	http.HandleFunc("/total", func(w http.ResponseWriter, _ *http.Request) {
		getTotalUseCase := usecase.NewGetTotalUseCase(repository)
		getTotalOutput, err := getTotalUseCase.Execute()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(getTotalOutput)
	})
	go http.ListenAndServe(":3333", nil)

	rabbitmqChannel, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer rabbitmqChannel.Close()

	messageDeliveryChannel := make(chan amqp.Delivery, 1000)

	go rabbitmq.Consume(rabbitmqChannel, messageDeliveryChannel)

	wg.Add(maxWorkers)

	for workerId := 1; workerId <= 3; workerId++ {
		defer wg.Done()
		go worker(workerId, messageDeliveryChannel, calculateFinalPriceUseCase)
	}

	wg.Wait()
}
