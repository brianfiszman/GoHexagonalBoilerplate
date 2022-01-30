package services

import "github.com/brianfiszman/GoHexagonalBoilerplate/pkg/domain/interfaces"

type HealthService struct {
	Database interfaces.Database
}

func (h HealthService) GetHealthiness() error {
	return h.Database.Ping()
}
