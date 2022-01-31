package http_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/application/containers/mocks"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/application/dtos"
)

func TestE2EHealth(t *testing.T) {
	var app mocks.MockAppContainer = mocks.NewMockAppContainer()

	handler := app.ServerContainer.Server.Router.Handler

	rr := httptest.NewRecorder()

	t.Run("should return 200", func(t *testing.T) {
		want := dtos.Response{
			Status:  http.StatusOK,
			Data:    nil,
			Message: dtos.HELLO_MSG,
		}

		req, err := http.NewRequest(http.MethodGet, "/health", rr.Body)

		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("HTTP Response Code: %v, Expected: %v", rr.Code, http.StatusOK)
		}

		res, err := serializeHealthResponse(rr.Body)

		if err != nil {
			t.Errorf("Error serializing response: %v", res)
		}

		if !reflect.DeepEqual(res, want) {
			t.Errorf("Results: %v and wanted %v are unequal", res, want)
		}
	})
}

func serializeHealthResponse(b *bytes.Buffer) (dtos.Response, error) {
	res := dtos.Response{}
	bodyBuffer, err := ioutil.ReadAll(b)

	if err != nil {
		return res, err
	}

	err = json.Unmarshal(bodyBuffer, &res)

	if err != nil {
		return res, err
	}

	return res, err
}
