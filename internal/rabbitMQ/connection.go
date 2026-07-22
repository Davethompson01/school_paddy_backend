package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}


func New(url string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        conn.Close()
        return nil, err
    }

    return &RabbitMQ{
        Conn: conn,
        Channel: ch,
    }, nil
}

func (r *RabbitMQ) Close() {
    r.Channel.Close()
    r.Conn.Close()
}