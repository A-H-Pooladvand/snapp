package configs

import (
	"github.com/caarlos0/env/v6"
)

type App struct {
	Port string `env:"APP_PORT"`
}

func NewApp() App {
	c := App{}
	if err := env.Parse(&c); err != nil {
		panic(err)
	}

	return c
}
