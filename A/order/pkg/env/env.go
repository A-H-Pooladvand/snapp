package env

import (
	"github.com/caarlos0/env/v6"
)

func Parse(val any) {
	if err := env.Parse(val); err != nil {
		panic(err)
	}
}
