package reciever

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Connection *amqp.Connection
var ConnectionError error

var GlobalExchange error

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Dialing the server to RabbitMQ
func DialToServer() {
	Connection, ConnectionError = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(ConnectionError, "Failed to connect to RabbitMQ server ")
}

// Finishing the Dial Server
func UndialServer() {
	Connection.Close()
}

// Creating channel to run processes
func CreateChannel() *amqp.Channel {
	channel, err := Connection.Channel()
	failOnError(err, "Failed to open a channel")
	return channel
}

func DeclareExchange() {
	channel := CreateChannel()
	GlobalExchange = channel.ExchangeDeclare(
		Exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(GlobalExchange, "Failed to create an exchange")
}

func (rc RecieverClient) Declare_Bind_Consume() {
	channel := CreateChannel()
	queue, err := channel.QueueDeclare(
		rc.Username, // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to Declare queue")

	err = channel.QueueBind(
		queue.Name, // queue name
		queue.Name, // routing key
		Exchange,   // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register consumer")
}

func (rc RecieverClient) recieve_log() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		Exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		rc.Username, // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,   // queue name
		q.Name,   // routing key
		Exchange, // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
