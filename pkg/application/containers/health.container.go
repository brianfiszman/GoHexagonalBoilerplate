package containers

import "github.com/brianfiszman/GoFromZeroToHero/pkg/domain/interfaces"

type HealthContainer struct {
	Database interfaces.Database
}

func NewHealthContainer() (h HealthContainer) {
	return
}
