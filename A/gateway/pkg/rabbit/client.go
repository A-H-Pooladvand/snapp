package rabbit

import (
	"gateway/configs"
	"github.com/streadway/amqp"
)

func Dial() (*amqp.Connection, error) {
	config := configs.NewRabbit()

	return amqp.Dial(config.URL())
}
