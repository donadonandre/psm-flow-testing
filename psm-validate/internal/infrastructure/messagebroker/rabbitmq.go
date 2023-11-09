package messagebroker

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	conn    *amqp091.Connection
	channel *amqp091.Channel
)

func InitRabbitMQ() *amqp091.Channel {
	conn, err := amqp091.Dial("amqp://user:default@localhost:5672/")
	if err != nil {
		log.Println("Error trying connect to RabbitMQ", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Println("Error creating Rabbit Channel", err)
	}

	return channel
}

func CloseRabbitMQ() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}
