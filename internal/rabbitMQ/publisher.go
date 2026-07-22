package rabbitmq

import (
	"encoding/json"

	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	"github.com/rabbitmq/amqp091-go"
)

func PublishBidCreated(ch *amqp091.Channel, apply solutionexpert_model.BidCreatedNotification) error {
	body, _ := json.Marshal(apply)
	return ch.Publish("", "Bid_created", true, true, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
