package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Rabbit struct {
	Host     string `env:"RABBITMQ_HOST"`
	Port     string `env:"RABBITMQ_PORT"`
	Username string `env:"RABBITMQ_USERNAME"`
	Password string `env:"RABBITMQ_PASSWORD"`
}

func NewRabbit() Rabbit {
	c := Rabbit{}
	if err := env.Parse(&c); err != nil {
		panic(err)
	}

	return c
}

func (r Rabbit) HasCredential() bool {
	// We can have a client with username but with empty password string
	// So we won't check for password in here, only username
	return r.Username != ""
}

func (r Rabbit) URL() string {
	if r.HasCredential() {
		return fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			r.Username,
			r.Password,
			r.Host,
			r.Port,
		)
	}

	return fmt.Sprintf("amqp://%s:%s/", r.Host, r.Port)
}
