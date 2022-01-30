package adapters_test

import (
	"testing"

	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/infrastructure/adapters/mocks"
)

func Test(t *testing.T) {
	pgxconn := mocks.NewMockPostgreSQLAdapter()

	t.Run("Should ping the database and fail", func(t *testing.T) {
		err := pgxconn.Ping()

		if err == nil {
			t.Error("Shouldn't ping if its disconnected")
		}
	})

	t.Run("Should ping the database", func(t *testing.T) {
		// Connect to database
		pgxconn.ConnectDatabase()

		err := pgxconn.Ping()

		if err != nil {
			t.Error(err)
		}
	})
}
