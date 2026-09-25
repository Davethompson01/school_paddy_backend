package main

import (
	"log"
	"os"

	"github.com/Davethompson01/School_Paddy_golang/database"
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	rabbitmq "github.com/Davethompson01/School_Paddy_golang/internal/rabbitMQ"
	worker "github.com/Davethompson01/School_Paddy_golang/internal/workers"
)

func main() {

	conn, err := database.DatabaseConnection()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer conn.Close()

	rabbit, err := rabbitmq.New(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal("RabbitMQ connection failed:", err)
	}
	defer rabbit.Close()

	cfg := config.ApiConfig{
		DB:     conn,
		Rabbit: rabbit,
	}

	log.Println("Notification worker started")

	worker.StartNotificationWorker(
		cfg.Rabbit.Channel,
		&cfg,
	)
}
