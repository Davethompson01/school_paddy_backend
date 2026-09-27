package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func ApplyBidCreated(ch *amqp.Channel) (<-chan amqp.Delivery, error) {

	return ch.Consume(
		"Bid_created",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func Negotiate_Bid(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		"Negotiate trasactions",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}


