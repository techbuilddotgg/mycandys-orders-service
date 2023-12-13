package rabbitmq

import (
	"context"
	"github.com/mycandys/orders/internal/env"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var Ch *amqp.Channel

func Connect() *amqp.Connection {
	url, _ := env.GetEnvVar(env.RABBITMQ_URL)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)

	}

	Ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Error opening RabbitMQ channel: %v", err)
	}

	return conn
}

func Close(conn *amqp.Connection) {
	if err := conn.Close(); err != nil {
		log.Fatalf("Error closing RabbitMQ connection: %v", err)
	}
}

func DeclareQueue(name string) amqp.Queue {
	q, err := Ch.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declaring RabbitMQ queue: %v", err)
	}
	return q
}

func Publish(exchangeName string, queueName string, body []byte, ctx context.Context) {
	err := Ch.PublishWithContext(ctx, exchangeName, queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		log.Fatalf("Error publishing message to RabbitMQ: %v", err)
	}

}
