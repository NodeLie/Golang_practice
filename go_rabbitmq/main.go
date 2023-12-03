package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Соединение с RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Создание канала
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Объявление очереди
	q, err := ch.QueueDeclare(
		"hello", // Имя очереди
		false,   // Durable
		false,   // Delete when unused
		false,   // Exclusive
		false,   // No-wait
		nil,     // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Отправка сообщения
	body := "Hello World!"
	err = ch.Publish(
		"",     // Exchange
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	fmt.Println(" [x] Sent", body)
}
