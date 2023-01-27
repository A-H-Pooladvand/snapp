package configs

import "order/pkg/env"

type Mysql struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	DATABASE string `env:"DB_DATABASE"`
	User     string `env:"DB_USERNAME"`
	Pass     string `env:"DB_PASSWORD"`
}

func NewMysql() Mysql {
	c := Mysql{}
	env.Parse(&c)

	return c
}
