package services

import "github.com/brianfiszman/GoFromZeroToHero/pkg/domain/interfaces"

type HealthService struct {
	Database interfaces.Database
}

func (h HealthService) GetHealthiness() error {
	return h.Database.Ping()
}
