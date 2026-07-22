package config

import (
	"database/sql"

	rabbitmq "github.com/Davethompson01/School_Paddy_golang/internal/rabbitMQ"
)

type ApiConfig struct {
	DB     *sql.DB
	Rabbit *rabbitmq.RabbitMQ
}
