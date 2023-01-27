package configs

import (
	"fmt"
	"order/pkg/env"
)

type App struct {
	Port string `env:"APP_PORT"`
	Mode string `env:"APP_ENV" envDefault:"release"`
}

func (a App) Addr() string {
	return fmt.Sprintf(":%s", a.Port)
}

func NewApp() App {
	c := App{}

	env.Parse(&c)

	return c
}
