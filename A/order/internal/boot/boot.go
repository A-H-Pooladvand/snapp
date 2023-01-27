package boot

import (
	"order/internal/interfaces"
)

var bootstraps = []interfaces.Booter{
	new(Env),
	new(Migrator),
}

type Bootstrap struct {
	Boots []interfaces.Booter
}

func (b *Bootstrap) Boot() {
	for _, boot := range b.Boots {
		if err := boot.Boot(); err != nil {
			//log.Error(err)
		}
	}
}

func New() *Bootstrap {
	return &Bootstrap{
		Boots: bootstraps,
	}
}
