package worker

import (
	"encoding/json"
	"log"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	rabbitmq "github.com/Davethompson01/School_Paddy_golang/internal/rabbitMQ"
	"github.com/Davethompson01/School_Paddy_golang/internal/respositary"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StartNotificationWorker(
	ch *amqp.Channel,
	api *config.ApiConfig,
) {

	msgs, err := rabbitmq.ApplyBidCreated(ch)
	if err != nil {
		log.Println("Failed to start notification worker:", err)
		return
	}

	log.Println("Notification worker started")

	for msg := range msgs {

		var event solutionexpert_model.BidCreatedNotification

		// Convert RabbitMQ JSON → Go struct
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Println("Invalid message:", err)

			msg.Nack(false, false)
			continue
		}

		log.Printf(
			"Processing bid notification: project=%d expert=%d student=%d bid=%d",

			event.StudentID,
			event.SolutionExpertID,
			event.ProjectID,
			event.BidID,
		)

		// Create notification in database
		event.Applied = true

		err := respositary.ApplyBidNotification(
			api,
			event,
		)

		if err != nil {
			log.Println("Failed to create notification:", err)

			// Tell RabbitMQ the message wasn't successfully processed
			msg.Nack(false, true)

			continue
		}

		// Everything succeeded
		err = msg.Ack(false)
		if err != nil {
			log.Println("Failed to ACK message:", err)
		}
	}

	log.Println("Notification worker stopped")
}
