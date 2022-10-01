<p align="center">
  <img
    height="300"
    src="./.github/golang-and-rabbitmq-banner.png"
    alt="Golang + Rabbitmq"
  />
</p>

# Golang + Rabbitmq example

## Overview

This application sends and consumes messages using Message Broker Rabbitmq, along with the power that Golang gives us to deal with situations that we need high performance.

- [x] Rabbitmq to Produce and Consume Messages
- [x] Concurrency using Goroutines
- [x] Channels
- [x] SQLite Database
- [x] HTTP Server
- [x] Workers
- [x] Docker + Docker Compose

---

<details>
<summary>
  Setup Application
</summary>

### Install dependencies

```bash
$ go mod tidy
```

### Create database

The database is already created and available at `internal/order/infra/database/sqlite.db`, but if you want to create it from scratch, just follow these steps:

```bash
$ cd internal/order/infra/database
$ touch sqlite.db
```

Now you need to access the database and create the `orders` table:

```sql
CREATE TABLE orders (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  price FLOAT NOT NULL,
  tax FLOAT NULL,
  final_price FLOAT NOT NULL
)
```

### Rabbitmq

#### Create orders queue

1. Go to http://localhost:15672 and log in to the Rabbitmq Management using guest credentials, by default the user is `guest` and the password is `guest`.
2. Choose the `Queues` tab.
3. Expand the `Add a new queue` option and set the queue name to `orders`.
4. Click on `Add queue`.

![rabbitmq-create-queue](.github/rabbitmq-create-queue.png)

#### Create bind

1. Choose the `Exchanges` tab.
2. Click on `amq.direct`.

![rabbitmq-exchange-direct](.github/rabbitmq-exchange-direct.png)

3. Expand the `Bindings` option and define in `To queue` the name `orders`.
4. Click on `Bind`.

![rabbitmq-create-exchange-bind](.github/rabbitmq-create-exchange-bind.png)

After creating the bind it will look like this

![rabbitmq-exchange-bind-complete](.github/rabbitmq-exchange-bind-complete.png)

### Produce messages

For you to be able to test it, just produce some messages using the producer, it is configured to generate 1.000 messages that will be stored in the `orders` queue.

```bash
$ go run cmd/producer/main.go
```

### Consume messages

Now to finish, just consume these messages, this procedure will occur concurrently and all messages that are consumed will be stored in the SQLite database.

```bash
$ go run cmd/main.go
```

</details>
