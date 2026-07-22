package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	rabbitmq "github.com/Davethompson01/School_Paddy_golang/internal/rabbitMQ"
	"github.com/Davethompson01/School_Paddy_golang/internal/respositary"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StartNotificationWorker(ch *amqp.Channel, api *config.ApiConfig) {

	msgs, err := rabbitmq.ApplyBidCreated(ch)
	if err != nil {
		log.Println(err)
		return
	}

	for msg := range msgs {

		var bid solutionexpert_model.ApplyForHomeWork
		var event solutionexpert_model.BidCreatedNotification

		if err := json.Unmarshal(msg.Body, &bid); err != nil {
			msg.Nack(false, false)
			continue
		}

		log.Println("Creating notification...")

		// create notification
		event.Applied = true
		updateNotis := respositary.ApplyBidNotification(api, event)
		if updateNotis != nil {
			msg.Nack(false, false)
		}

		msg.Ack(false)

	}
}
