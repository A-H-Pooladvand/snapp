package boot

import (
	"gateway/internal/interfaces"
	"gateway/pkg/log"
)

var bootstraps = []interfaces.Booter{
	new(Env),
}

type Bootstrap struct {
	Boots []interfaces.Booter
}

func New() *Bootstrap {
	return &Bootstrap{
		Boots: bootstraps,
	}
}

func (b *Bootstrap) Boot() {
	for _, boot := range b.Boots {
		if err := boot.Boot(); err != nil {
			log.Error(err)
		}
	}
}
