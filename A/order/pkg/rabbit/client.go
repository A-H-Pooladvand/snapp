package rabbit

import (
	"github.com/streadway/amqp"
	"order/configs"
)

func Dial() (*amqp.Connection, error) {
	config := configs.NewRabbit()

	return amqp.Dial(config.URL())
}
