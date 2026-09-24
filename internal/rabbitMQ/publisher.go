package rabbitmq

import (
	"encoding/json"
	"errors"

	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	"github.com/rabbitmq/amqp091-go"
)

func PublishBidCreated(
	ch *amqp091.Channel,
	apply solutionexpert_model.BidCreatedNotification,
) error {

	if ch == nil {
		return errors.New("RabbitMQ channel is nil")
	}

	if ch.IsClosed() {
		return errors.New("RabbitMQ channel is closed")
	}

	body, err := json.Marshal(apply)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		"Bid_created",
		true,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
