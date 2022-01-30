package mocks

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type MockPostgreSQLAdapter struct {
	isConnected bool
}

func NewMockPostgreSQLAdapter() (adapter *MockPostgreSQLAdapter) {
	adapter = &MockPostgreSQLAdapter{}

	return
}

func (d *MockPostgreSQLAdapter) ConnectDatabase() {
	d.isConnected = true

	logrus.Info("Connected to localhost:5432")
}

func (d *MockPostgreSQLAdapter) Ping() error {
	if !d.isConnected {
		return errors.New("Not connected to the database")
	}

	return nil
}
