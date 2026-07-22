package main

import (
	"log"
	"os"

	rabbitmq "github.com/Davethompson01/School_Paddy_golang/internal/rabbitMQ"
)

func main() {
	rabbit, err := rabbitmq.New(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()
}
